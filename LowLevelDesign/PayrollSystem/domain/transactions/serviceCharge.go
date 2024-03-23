package transactions

import (
	"context"
	"fmt"
	"payroll"
	"payroll/domain/affiliations"
	"payroll/domain/employee"
	"time"
)

type AddServiceChargeTransaction struct {
	memberId   string
	charge     float64
	recordDate time.Time
}

func NewAddServiceChargeTransaction(memberId string, charge float64, recordDate time.Time) payroll.Transaction {
	return &AddServiceChargeTransaction{
		memberId:   memberId,
		charge:     charge,
		recordDate: recordDate,
	}
}

func (a *AddServiceChargeTransaction) Execute(ctx context.Context) error {
	e := payroll.PayrollDatabase.GetUnionMember(a.memberId)
	if e == nil {
		return fmt.Errorf("member does not exist with id %s", a.memberId)
	}
	emp := e.(*employee.Employee)
	employeeAffiliations := emp.GetAffiliations(ctx)
	for _, affiliation := range employeeAffiliations {
		if unionAffiliation, ok := affiliation.(*affiliations.UnionAffiliation); ok {
			unionAffiliation.SubmitServiceCharge(context.Background(), a.charge, a.recordDate)
			return nil
		}
	}
	return fmt.Errorf("no union affiliation found for employee with member id %s", a.memberId)
}
