package command

import (
	MetroCard "Metro"
	"fmt"
	"log/slog"
)

type Command interface {
	Execute() Result
	BuildInternal()
}

type Result interface {
	IsSuccess() bool
	GetResult() any
}

type defaultResult struct {
	isSuccess bool
	result    any
}

func (r *defaultResult) IsSuccess() bool {
	return r.isSuccess
}

func (r *defaultResult) GetResult() any {
	return r.result
}

func NewSuccessResult(result any) Result {
	return &defaultResult{isSuccess: true, result: result}
}

var cardRepository *MetroCard.CardRepository
var Log *MetroCard.Log
var boardingManagerRepository map[string]MetroCard.BoardingManager

func Setup(logger *slog.Logger) {
	cardRepository = MetroCard.NewCardRepository()
	Log = MetroCard.NewLog()
	boardingManagerRepository = make(map[string]MetroCard.BoardingManager)
	boardingManagerRepository["CENTRAL"] = MetroCard.NewDefaultBoardingManager(MetroCard.NewDefaultServiceFee(),
		Log, MetroCard.NewStation("CENTRAL"), logger)
	boardingManagerRepository["AIRPORT"] = MetroCard.NewDefaultBoardingManager(MetroCard.NewDefaultServiceFee(),
		Log, MetroCard.NewStation("AIRPORT"), logger)
}

func Factory(command []string) Command {
	var processedCommand Command
	switch command[0] {
	case "BALANCE":
		processedCommand = &BalanceCommand{
			command: command[1:],
		}
	case "CHECK_IN":
		processedCommand = &CheckInCommand{
			command: command[1:],
		}
	case "PRINT_SUMMARY":
		processedCommand = &PrintSummaryCommand{}
	}
	processedCommand.BuildInternal()
	return processedCommand
}

func PrintSummary() {
	for _, val := range []string{"CENTRAL", "AIRPORT"} {
		collection, discount := Log.GetCollectionAndDiscountOfStation(MetroCard.NewStation(val))
		fmt.Printf("TOTAL_COLLECTION %s %d %d\n", val, int(collection), int(discount))
		fmt.Println("PASSENGER_TYPE_SUMMARY")
		for _, v := range Log.GetPassengerSummaryForStation(MetroCard.NewStation(val)) {
			fmt.Printf("%s %d\n", v.PassengerType, v.Count)
		}
	}
}
