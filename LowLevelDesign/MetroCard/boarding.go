package MetroCard

import (
	"log/slog"
)

type BoardingManager interface {
	// Board processes the journey
	Board(journey PassengerJourney)
}

type ServiceFee interface {
	GetServiceFee(amount float64) float64
}

type DefaultServiceFee struct {
}

func NewDefaultServiceFee() *DefaultServiceFee {
	return &DefaultServiceFee{}
}

func (s *DefaultServiceFee) GetServiceFee(amount float64) float64 {
	return amount * 0.02
}

type Collection struct {
	Card   MetroCard
	Amount float64
}

type DefaultBoardingManager struct {
	ServiceFee  ServiceFee
	Collections []Collection
	Station     Station
	Log         *Log
	logger      *slog.Logger
}

func NewDefaultBoardingManager(ServiceFee ServiceFee, log *Log, station Station, logger *slog.Logger) *DefaultBoardingManager {
	return &DefaultBoardingManager{
		ServiceFee: ServiceFee,
		Station:    station,
		Log:        log,
		logger:     logger,
	}
}

func (b *DefaultBoardingManager) Board(journey PassengerJourney) {
	var serviceFee float64
	fare, discount := journey.GetFareAndDiscount(*b.Log)
	b.logger.Info("Boarding", "fare", fare, "discount", discount, "journey", journey)
	if !b.IsBalanceSufficient(journey) {
		rechargeAmount := fare - journey.Card.GetBalance() - discount
		serviceFee = b.ServiceFee.GetServiceFee(rechargeAmount)
		journey.Card.Recharge(int(rechargeAmount + serviceFee))
	}
	journey.SetDiscount(discount)
	journey.DeductCharge(serviceFee + fare - discount)
	b.Log.AddJourney(&journey)
}

func (b *DefaultBoardingManager) IsBalanceSufficient(journey PassengerJourney) bool {
	fare, discount := journey.GetFareAndDiscount(*b.Log)
	return (fare - discount) < journey.Card.GetBalance()
}

// DeductMoney deducts specified amount from journey's card and returns remaining balance
func (b *DefaultBoardingManager) DeductMoney(journey PassengerJourney, amount float64) float64 {
	balance, _ := journey.DeductCharge(amount)
	return balance
}
