package affiliations

import (
	"context"
	"payroll/domain/payment"
	"time"
)

type AffiliationType string

type Affiliation interface {
	SubmitServiceCharge(ctx context.Context, serviceCharge float64, recordDate time.Time)
}

type UnionAffiliation struct {
	Name           string
	serviceCharges []payment.ServiceCharge
}

func (u *UnionAffiliation) SubmitServiceCharge(ctx context.Context, serviceCharge float64, recordDate time.Time) {
	u.serviceCharges = append(u.serviceCharges, payment.NewServiceCharge(serviceCharge, recordDate))
}

func NewUnionAffiliation(name string) Affiliation {
	return &UnionAffiliation{
		Name: name,
	}
}
