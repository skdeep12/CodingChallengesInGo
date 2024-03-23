package transactions

import (
	"context"
	"payroll"
)

type DeleteEmployeeTransasction struct {
	empID string
}

func NewDeleteEmployeeTransaction(empID string) payroll.Transaction {
	return &DeleteEmployeeTransasction{
		empID: empID,
	}
}

func (d *DeleteEmployeeTransasction) Execute(ctx context.Context) error {
	payroll.PayrollDatabase.DeleteEmployee(d.empID)
	return nil
}
