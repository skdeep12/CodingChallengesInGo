package ParkingLot

type ParkingFeeStrategy interface {
	CalculateCost(hours int) int
}

func ParkingFeeFactory(vehicleType VehicleType) ParkingFeeStrategy {
	switch vehicleType {
	case BikeVehicleType:
		return NewBikeFeeStrategy(5)
	case CarVehicleType:
		return NewCompactCarStrategy(10)
	}
	return nil
}

type BikeFeeStrategy struct {
	chargesPerHour int
}

func (b *BikeFeeStrategy) CalculateCost(hours int) int {
	return b.chargesPerHour * hours
}

func NewBikeFeeStrategy(chargesPerHour int) *BikeFeeStrategy {
	return &BikeFeeStrategy{
		chargesPerHour: chargesPerHour,
	}
}

type CompactCarStrategy struct {
	chargesPerHour int
}

func (c CompactCarStrategy) CalculateCost(hours int) int {
	return c.chargesPerHour * hours
}

func NewCompactCarStrategy(chargesPerHour int) CompactCarStrategy {
	return CompactCarStrategy{
		chargesPerHour: chargesPerHour,
	}
}
