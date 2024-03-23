package transactions

import (
	"context"
	"payroll"
	"payroll/domain/employee"
	"payroll/domain/payment"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSalariedEmployee(t *testing.T) {
	// add employee transaction, default should be
	txn := NewAddSalariedEmployeeTransaction("123", "Subhash", "Home", 100)
	txn.Execute(context.Background())
	emp := payroll.PayrollDatabase.GetEmployee("123").(*employee.Employee)
	classification := emp.GetClassification(context.Background())
	_, ok := classification.(*payment.SalariedClassification)
	assert.Equal(t, ok, true)
	schedule := emp.GetPaymentSchedule(context.Background())
	_, scheduleOk := schedule.(*payment.MonthlySchedule)
	assert.Equal(t, scheduleOk, true)
	method := emp.GetPaymentMethod(context.Background())
	_, methodOk := method.(*payment.HoldPaymentMethod)
	assert.Equal(t, methodOk, true)
}

func TestCommissionedEmployee(t *testing.T) {
	txn := NewCommissionedEmployeeTransaction("1234", "Sunitq", "Home", 100, 2)
	txn.Execute(context.Background())
	emp := payroll.PayrollDatabase.GetEmployee("1234").(*employee.Employee)
	classification := emp.GetClassification(context.Background())
	_, ok := classification.(*payment.CommissionedClassfication)
	assert.Equal(t, ok, true)
	schedule := emp.GetPaymentSchedule(context.Background())
	_, scheduleOk := schedule.(*payment.BiWeeklySchedule)
	assert.Equal(t, scheduleOk, true)
	method := emp.GetPaymentMethod(context.Background())
	_, methodOk := method.(*payment.HoldPaymentMethod)
	assert.Equal(t, methodOk, true)
}

func TestDeleteEmployee(t *testing.T) {
	txn := NewCommissionedEmployeeTransaction("1234", "Sunitq", "Home", 100, 2)
	txn.Execute(context.Background())
	emp := payroll.PayrollDatabase.GetEmployee("1234").(*employee.Employee)
	assert.NotNil(t, emp)
	deleteTxn := NewDeleteEmployeeTransaction("1234")
	deleteTxn.Execute(context.Background())
	deletedEmployee := payroll.PayrollDatabase.GetEmployee("1234")
	assert.Nil(t, deletedEmployee)
}

func TestAddTimeCardTransaction(t *testing.T) {
	startTime := time.Now()
	employeeTxn := NewAddHourlyEmployeeTransaction("1234", "Sunitq", "Home", 100)
	employeeTxn.Execute(context.Background())
	txn := NewTimeCardTransaction("1234", time.Now(), 10)
	txn.Execute(context.Background())
	emp := payroll.PayrollDatabase.GetEmployee("1234").(*employee.Employee)
	classification := emp.GetClassification(context.Background())
	hourlyClassification, ok := classification.(*payment.HourlyClassification)
	assert.Equal(t, ok, true)
	timeCards := hourlyClassification.GetTimeCards(context.Background(), startTime, time.Now())
	assert.Equal(t, len(timeCards), 1)
}

func TestAddServiceChargeTransaction_Execute(t *testing.T) {
	employeeTxn := NewAddHourlyEmployeeTransaction("1234", "Sunitq", "Home", 100)
	employeeTxn.Execute(context.Background())
	unionTxn := NewAddUnionAffiliationTransaction("1234", "58")
	unionTxn.Execute(context.Background())
	txn := NewAddServiceChargeTransaction("58", 10, time.Now())
	txn.Execute(context.Background())
	emp := payroll.PayrollDatabase.GetEmployee("1234").(*employee.Employee)
	emp.GetAffiliation(context.Background())
	serviceCharges := emp
	assert.Equal(t, len(serviceCharges), 1)
}
