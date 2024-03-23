package payment

import (
	"context"
	"fmt"
)

type Method interface {
	ExecutePayment(ctx context.Context, payment float64)
}

type HoldPaymentMethod struct {
}

func NewHoldPaymentMethod() Method {
	return &HoldPaymentMethod{}
}

func (h *HoldPaymentMethod) ExecutePayment(ctx context.Context, payment float64) {
	fmt.Printf("hold payment to postmaster, %.2f", payment)
}
