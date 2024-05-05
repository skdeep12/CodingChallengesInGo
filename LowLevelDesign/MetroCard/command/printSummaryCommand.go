package command

type PrintSummaryCommand struct {
}

func (p *PrintSummaryCommand) Execute() Result {
	PrintSummary()
	return NewSuccessResult(nil)
}

func (p *PrintSummaryCommand) BuildInternal() {
	return
}
