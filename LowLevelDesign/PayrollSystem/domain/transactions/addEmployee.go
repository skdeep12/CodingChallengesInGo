package transactions

import (
	"context"
)

type AddEmployeeTransaction struct {
	ID        string
	Address   string
	FirstName string
	LastName  string
}

func (a *AddEmployeeTransaction) Execute(ctx context.Context) error {
	panic("implement me")
	// add default classification and payment method, payment schedule
}
