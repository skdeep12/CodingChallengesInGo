package transactions

import (
	"context"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
)

type AddHourlyEmployeeTransaction struct {
	AddEmployeeTransaction
	HourlyRate float64
}

func NewAddHourlyEmployeeTransaction(id, name, address string, hourlyRate float64) payroll.Transaction {
	return &AddHourlyEmployeeTransaction{
		HourlyRate: hourlyRate,
		AddEmployeeTransaction: AddEmployeeTransaction{
			ID:        id,
			FirstName: name,
			Address:   address,
		},
	}
}

func (a *AddHourlyEmployeeTransaction) Execute(ctx context.Context) error {
	e := employee.NewEmployee(a.ID, a.FirstName+" "+a.LastName, a.Address)
	e.SetClasstification(ctx, payment.NewHourlyEmployeeClassification(a.HourlyRate))
	e.SetPaymentMethod(ctx, payment.NewHoldPaymentMethod())
	e.SetPaymentSchedule(ctx, payment.NewWeeklySchedule())
	payroll.PayrollDatabase.AddEmployee(&e)
	return nil
}
