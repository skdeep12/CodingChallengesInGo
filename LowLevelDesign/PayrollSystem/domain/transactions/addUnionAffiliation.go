package transactions

import (
	"context"
	"payroll"
	"payroll/domain/affiliations"
	"payroll/domain/employee"
)

type AddUnionAffiliationTransaction struct {
	empId    string
	memberId string
}

func NewAddUnionAffiliationTransaction(empId string, memberId string) payroll.Transaction {
	return &AddUnionAffiliationTransaction{
		empId:    empId,
		memberId: memberId,
	}
}

func (a *AddUnionAffiliationTransaction) Execute(ctx context.Context) error {
	if err := payroll.PayrollDatabase.AddUnionMember(a.memberId, a.empId); err != nil {
		return err
	}
	payroll.PayrollDatabase.GetEmployee(a.empId).(*employee.Employee).SetAffiliation(ctx, affiliations.NewUnionAffiliation(a.memberId))
	return nil
}
