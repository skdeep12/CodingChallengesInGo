package main

import (
	"Metro/command"
	"bufio"
	"log/slog"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		println("Please provide an input file")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	file, err := os.Open(inputFile)
	if err != nil {
		println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	command.Setup(logger)
	for scanner.Scan() {
		commandText := scanner.Text()
		inputcommand := strings.Split(commandText, " ")
		processedCommand := command.Factory(inputcommand)
		processedCommand.Execute()
		//command.PrintSummary()
		//println("=============")
	}
}
