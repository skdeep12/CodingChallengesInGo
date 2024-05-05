package command

import (
	"Metro"
	"strconv"
)

type BalanceCommand struct {
	command []string
	card    MetroCard.MetroCard
}

func NewBalanceCommand(command []string) *BalanceCommand {
	return &BalanceCommand{
		command: command,
	}
}

func (b *BalanceCommand) BuildInternal() {
	balance, _ := strconv.Atoi(b.command[1])
	b.card = MetroCard.NewCard(b.command[0], float64(balance))
}

func (b *BalanceCommand) Execute() Result {
	cardRepository.AddCard(b.card)
	return NewSuccessResult(nil)
}
