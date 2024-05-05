package command

import (
	MetroCard "Metro"
	"testing"
)

func TestBalanceCommand_Execute(t *testing.T) {
	cardRepository = MetroCard.NewCardRepository()
	processedCommand1 := Factory([]string{"BALANCE", "AB123", "100"})
	processedCommand2 := Factory([]string{"BALANCE", "AB1234", "150"})
	processedCommand1.Execute()
	processedCommand2.Execute()
	PrintSummary()
}
