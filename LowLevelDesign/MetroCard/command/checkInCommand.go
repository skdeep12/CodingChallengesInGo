package command

import MetroCard "Metro"

type CheckInCommand struct {
	command []string
	journey *MetroCard.PassengerJourney
}

func NewCheckInCommand(command []string) *CheckInCommand {
	return &CheckInCommand{
		command: command,
	}
}

// BuildInternal builds the internal representation of the command
// AB123 ADULT AIRPORT
func (c *CheckInCommand) BuildInternal() {
	card := cardRepository.GetCard(c.command[0])
	c.journey = MetroCard.NewJourney(card, c.command[1], c.command[2])
}

func (c *CheckInCommand) Execute() Result {
	boardingManagerRepository[c.command[2]].Board(*c.journey)
	return NewSuccessResult(nil)
}
