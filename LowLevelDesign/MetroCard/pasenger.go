package MetroCard

type PassengerType string

const (
	Kid           PassengerType = "KID"
	Adult         PassengerType = "ADULT"
	SeniorCitizen PassengerType = "SENIOR_CITIZEN"
)

type PassengerFare interface {
	GetFare() float64
	GetType() PassengerType
}

func NewPassengerFare(passengerType PassengerType) PassengerFare {
	switch passengerType {
	case Kid:
		return NewKidPassengerFare()
	case Adult:
		return NewAdultPassengerFare()
	case SeniorCitizen:
		return NewSeniorCitizenPassengerFare()
	}
	return nil
}

type KidPassengerFare struct {
	passengerType PassengerType
}

func (k *KidPassengerFare) GetType() PassengerType {
	return k.passengerType
}

func NewKidPassengerFare() *KidPassengerFare {
	return &KidPassengerFare{passengerType: Kid}
}

type AdultPassengerFare struct {
	passengerType PassengerType
}

func (a *AdultPassengerFare) GetType() PassengerType {
	return a.passengerType
}

func NewAdultPassengerFare() *AdultPassengerFare {
	return &AdultPassengerFare{passengerType: Adult}
}

type SeniorCitizenPassengerFare struct {
	passengerType PassengerType
}

func (s *SeniorCitizenPassengerFare) GetType() PassengerType {
	return s.passengerType
}

func NewSeniorCitizenPassengerFare() *SeniorCitizenPassengerFare {
	return &SeniorCitizenPassengerFare{passengerType: SeniorCitizen}
}

func (k *KidPassengerFare) GetFare() float64 {
	return 50
}

func (a *AdultPassengerFare) GetFare() float64 {
	return 200
}

func (s *SeniorCitizenPassengerFare) GetFare() float64 {
	return 100
}
