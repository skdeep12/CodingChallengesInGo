package payment

import "time"

type ServiceCharge struct {
	Charge     float64
	RecordDate time.Time
}

func NewServiceCharge(charge float64, recordDate time.Time) ServiceCharge {
	return ServiceCharge{
		Charge:     charge,
		RecordDate: recordDate,
	}
}
