package payroll

import "context"

type Transaction interface {
	Execute(ctx context.Context) error
}
