package payment

import (
	"context"
	"time"
)

type Classification interface {
	CalculatePayout(context.Context) float64
}

type SalariedClassification struct {
	salary float64
}

func NewSalariedEmployeeClassification(salary float64) Classification {
	return &SalariedClassification{
		salary: salary,
	}
}

func (s *SalariedClassification) CalculatePayout(ctx context.Context) float64 {
	return s.salary
}

type HourlyClassification struct {
	hourlyRate float64
	timeCards  []TimeCard
}

func NewHourlyEmployeeClassification(hourlyRate float64) Classification {
	return &HourlyClassification{
		hourlyRate: hourlyRate,
	}
}

func (h *HourlyClassification) CalculatePayout(ctx context.Context) float64 {
	// salary + commision
	return h.hourlyRate
}

func (h *HourlyClassification) AddTimeCard(ctx context.Context, timeCard TimeCard) {
	h.timeCards = append(h.timeCards, timeCard)
}

func (h *HourlyClassification) GetTimeCards(ctx context.Context, startDate, endDate time.Time) []TimeCard {
	var timeCards []TimeCard
	for _, timeCard := range h.timeCards {
		if timeCard.CreatedAt.Before(startDate) || timeCard.CreatedAt.After(endDate) {
			continue
		} else {
			timeCards = append(timeCards, timeCard)
		}
	}
	return timeCards
}

type CommissionedClassfication struct {
	salary         float64
	commissionRate float64
	saleReceipts   []SaleReceipt
}

func NewCommissionedClassification(salary, commissionRate float64) Classification {
	return &CommissionedClassfication{
		salary:         salary,
		commissionRate: commissionRate,
	}
}
func (h *CommissionedClassfication) CalculatePayout(ctx context.Context) float64 {
	// salary + commision
	return h.salary
}

func (s *CommissionedClassfication) SubmitSaleReceipt(ctx context.Context, saleReceipt SaleReceipt) {
	s.saleReceipts = append(s.saleReceipts, saleReceipt)
}
