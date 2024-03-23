package transactions

import (
	"context"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
)

type AddSalariedEmployeeTransaction struct {
	AddEmployeeTransaction
	Salary float64
}

func NewAddSalariedEmployeeTransaction(id, name, address string, salary float64) payroll.Transaction {
	return &AddSalariedEmployeeTransaction{
		Salary: salary,
		AddEmployeeTransaction: AddEmployeeTransaction{
			ID:        id,
			FirstName: name,
			Address:   address,
		},
	}
}

func (a *AddSalariedEmployeeTransaction) Execute(ctx context.Context) error {
	e := employee.NewEmployee(a.ID, a.FirstName+" "+a.LastName, a.Address)
	e.SetClasstification(ctx, payment.NewSalariedEmployeeClassification(a.Salary))
	e.SetPaymentMethod(ctx, payment.NewHoldPaymentMethod())
	e.SetPaymentSchedule(ctx, payment.NewMonthlySchedule())
	payroll.PayrollDatabase.AddEmployee(&e)
	return nil
}
