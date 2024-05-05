package command

import (
	"log/slog"
	"os"
	"testing"
)

func TestCheckInCommand_Execute(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	Setup(logger)
	processedCommand1 := Factory([]string{"BALANCE", "AB123", "100"})
	processedCommand2 := Factory([]string{"BALANCE", "AB1234", "150"})
	processedCommand1.Execute()
	processedCommand2.Execute()
	//PrintSummary()
	processedCommand := Factory([]string{"CHECK_IN", "AB123", "ADULT", "CENTRAL"})
	processedCommand.Execute()
	PrintSummary()
}
