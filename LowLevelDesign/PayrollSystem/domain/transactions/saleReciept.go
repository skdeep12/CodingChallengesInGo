package transactions

import (
	"context"
	"fmt"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
	"time"
)

type AddSaleReceiptTransaction struct {
	empId      string
	saleAmount float64
	recordDate time.Time
}

func NewAddSaleReceiptTransaction(empId string, saleAmount float64, recordDate time.Time) *AddSaleReceiptTransaction {
	return &AddSaleReceiptTransaction{
		empId:      empId,
		saleAmount: saleAmount,
		recordDate: recordDate,
	}
}

func (a *AddSaleReceiptTransaction) Execute(ctx context.Context) error {
	e := payroll.PayrollDatabase.GetEmployee(a.empId)
	if e == nil {
		return fmt.Errorf("employee does not exist with id %s", a.empId)
	}
	emp := e.(*employee.Employee)
	paymentClassification := emp.GetClassification(context.Background())
	commissionedClassification, ok := paymentClassification.(*payment.CommissionedClassfication)
	if !ok {
		return fmt.Errorf("employee is not commissioned employee")
	}
	commissionedClassification.SubmitSaleReceipt(context.Background(), payment.NewSaleReceipt(a.saleAmount, a.recordDate))
	return nil
}
