package transactions

import (
	"context"
	"fmt"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
	"time"
)

type TimeCardTransaction struct {
	empID string
	date  time.Time
	hours int
}

func NewTimeCardTransaction(empID string, date time.Time, hours int) payroll.Transaction {
	return &TimeCardTransaction{
		empID: empID,
		date:  date,
		hours: hours,
	}
}

func (t *TimeCardTransaction) Execute(ctx context.Context) error {
	e := payroll.PayrollDatabase.GetEmployee(t.empID)
	if e == nil {
		return fmt.Errorf("employee does not exist with id %s", t.empID)
	}
	emp := e.(*employee.Employee)
	paymentClassification := emp.GetClassification(context.Background())
	hourlyClassification, hourlyOk := paymentClassification.(*payment.HourlyClassification)
	if !hourlyOk {
		return fmt.Errorf("employee is not hourly employee")
	}
	hourlyClassification.AddTimeCard(context.Background(), payment.NewTimeCard(t.hours, t.date))
	return nil
}
