package ParkingLot

import "time"

type VehicleType string

const (
	CarVehicleType  VehicleType = "car"
	BikeVehicleType VehicleType = "bike"
)

type Vehicle struct {
	registrationNumber string
	feeStrategy        ParkingFeeStrategy
	parkedAt           time.Time
	exitedAt           time.Time
	parkingSlot        *ParkingSlot
}

func (v *Vehicle) GetParkingSlot() *ParkingSlot {
	return v.parkingSlot
}

func (v *Vehicle) CalculateCharges(exitAt time.Time) int {
	diff := exitAt.Sub(v.parkedAt)
	hours := int(diff.Hours())
	if float64(hours) < diff.Hours() {
		hours++
	}
	return v.feeStrategy.CalculateCost(hours)
}

func (v *Vehicle) SetParkingSlot(slot *ParkingSlot) {
	v.parkingSlot = slot
}

func NewVehicle(registrationNumber string, vehicleType VehicleType) Vehicle {
	return Vehicle{
		registrationNumber: registrationNumber,
		feeStrategy:        ParkingFeeFactory(vehicleType),
	}
}
