package transactions

import (
	"context"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
)

type AddCommissionedEmployeeTransaction struct {
	AddEmployeeTransaction
	Salary         float64
	CommissionRate float64
}

func NewCommissionedEmployeeTransaction(id, name, address string, salary, commissionRate float64) payroll.Transaction {
	return &AddCommissionedEmployeeTransaction{
		Salary:         salary,
		CommissionRate: commissionRate,
		AddEmployeeTransaction: AddEmployeeTransaction{
			ID:        id,
			FirstName: name,
			Address:   address,
		},
	}
}

func (a *AddCommissionedEmployeeTransaction) Execute(ctx context.Context) error {
	e := employee.NewEmployee(a.ID, a.FirstName+" "+a.LastName, a.Address)
	e.SetClasstification(ctx, payment.NewCommissionedClassification(a.Salary, a.CommissionRate))
	e.SetPaymentMethod(ctx, payment.NewHoldPaymentMethod())
	e.SetPaymentSchedule(ctx, payment.NewBiWeeklySchedule())
	payroll.PayrollDatabase.AddEmployee(&e)
	return nil
}
