package employee

import (
	"context"
	"payroll/domain/affiliations"
	"payroll/domain/payment"
)

// Employee has EmployeeClassfication, Affiliation, Method
// EmployeeClassfication has rate/fixed and Schedule

type Employee struct {
	Name            string
	ID              string
	address         string
	classification  payment.Classification
	paymentMethod   payment.Method
	paymentSchedule payment.Schedule
	affiliations    affiliations.Affiliation
}

func NewEmployee(id, name string, address string) Employee {
	return Employee{
		Name:    name,
		address: address,
		ID:      id,
	}
}

func (e *Employee) SetClasstification(ctx context.Context, classification payment.Classification) {
	e.classification = classification
}

func (e *Employee) GetClassification(ctx context.Context) payment.Classification {
	return e.classification
}

func (e *Employee) SetPaymentMethod(ctx context.Context, pMethod payment.Method) {
	e.paymentMethod = pMethod
}

func (e *Employee) GetAffiliation(ctx context.Context) affiliations.Affiliation {
	return e.affiliations
}

func (e *Employee) GetPaymentMethod(ctx context.Context) payment.Method {
	return e.paymentMethod
}

func (e *Employee) SetPaymentSchedule(ctx context.Context, pSchedule payment.Schedule) {
	e.paymentSchedule = pSchedule
}

func (e *Employee) GetPaymentSchedule(ctx context.Context) payment.Schedule {
	return e.paymentSchedule
}

func (e *Employee) SetAffiliation(ctx context.Context, affiliation affiliations.Affiliation) {
	e.affiliations = affiliation
}

func (e *Employee) GetID() string {
	return e.ID
}
