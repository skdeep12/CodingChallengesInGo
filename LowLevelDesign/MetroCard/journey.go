package MetroCard

type PassengerJourney struct {
	Passenger  PassengerFare
	Start      Station
	End        Station
	Card       MetroCard
	Collection float64
	Discount   float64
}

func NewJourney(card MetroCard, passengerType string, start string) *PassengerJourney {
	var end string
	if start == "CENTRAL" {
		end = "AIRPORT"
	} else {
		end = "CENTRAL"
	}
	return &PassengerJourney{
		Card:      card,
		Passenger: NewPassengerFare(PassengerType(passengerType)),
		Start:     NewStation(start),
		End:       NewStation(end),
	}
}

func (j *PassengerJourney) IsBalanceSufficient(log Log) bool {
	// based on station calcuate charge
	fare, _ := j.GetFareAndDiscount(log)
	if fare > j.Card.GetBalance() {
		return false
	}
	return true
}

func (j *PassengerJourney) GetDeficientBalance(log Log) float64 {
	fare, _ := j.GetFareAndDiscount(log)
	if fare > j.Card.GetBalance() {
		return 0
	}
	return j.Card.GetBalance() - fare
}

func (j *PassengerJourney) DeductCharge(amount float64) (float64, error) {
	// based on station calcuate charge
	j.Collection = amount
	return j.Card.DeductCharge(amount)
}

// GetFareAndDiscount returns total fare and discount
func (j *PassengerJourney) GetFareAndDiscount(log Log) (float64, float64) {
	// based on stations this charge might vary
	charge := j.Passenger.GetFare()
	if log.IsReturnJourney(j) {
		charge /= 2
	}
	return j.Passenger.GetFare(), j.Passenger.GetFare() - charge
}

func (j *PassengerJourney) SetDiscount(discount float64) {
	j.Discount = discount
}

func (j *PassengerJourney) String() string {
	return j.Start.Name + " -> " + j.End.Name + " " + string(j.Passenger.GetType()) + " " + string(j.Card.GetNumber())
}
