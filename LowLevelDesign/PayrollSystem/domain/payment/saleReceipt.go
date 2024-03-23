package payment

import "time"

type SaleReceipt struct {
	SaleAmount float64
	RecordDate time.Time
}

func NewSaleReceipt(saleAmount float64, recordDate time.Time) SaleReceipt {
	return SaleReceipt{
		SaleAmount: saleAmount,
		RecordDate: recordDate,
	}
}
