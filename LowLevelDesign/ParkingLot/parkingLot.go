package ParkingLot

import (
	"fmt"
	"time"
)

type ParkingLotInterface interface {
	ParkVehicle(vehicle Vehicle) (*ParkingSlot, error)
	CalculateCost(vehicle Vehicle) int
	UnParkVehicle(vehicle Vehicle) *ParkingSlot
	GetAvailableSlot(vehicle Vehicle) *ParkingSlot
}

type Location struct {
	LatLng string
}

type ParkingLot struct {
	id       int
	location Location
	levels   []ParkingLevel
}

func (p *ParkingLot) GetAvailableSlot(vehicle Vehicle) *ParkingSlot {
	for _, level := range p.levels {
		if slot := level.GetAvailableParkingSlot(vehicle); slot != nil {
			return slot
		}
	}
	return nil
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle) (*ParkingSlot, error) {
	if slot := p.GetAvailableSlot(vehicle); slot != nil {
		slot.MarkOccupied()
		vehicle.SetParkingSlot(slot)
		return slot, nil
	}
	return nil, fmt.Errorf("no available slot found")
}

func (p *ParkingLot) UnParkVehicle(vehicle Vehicle) *ParkingSlot {
	charges := vehicle.CalculateCharges(time.Now())
	payCharges(vehicle, charges)
	slot := vehicle.GetParkingSlot()
	slot.MarkFree()
	return slot
}

func NewParkingLot(id int, location Location) ParkingLot {
	return ParkingLot{
		id:       id,
		location: location,
	}
}

func (p *ParkingLot) AddParkingLevel(level int, row, col int) {
	p.levels = append(p.levels, NewParkingLevel(level, row, col))
}

type ParkingLevel struct {
	id    int
	floor int
	slots []*ParkingSlot
}

func (p *ParkingLevel) GetAvailableParkingSlot(vehicle Vehicle) *ParkingSlot {
	for _, slot := range p.slots {
		if slot.IsFree() {
			return slot
		}
	}
	return nil
}
func NewParkingLevel(floor int, row, col int) ParkingLevel {
	slots := make([]*ParkingSlot, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			slots = append(slots, NewParkingSlot(i, j))
		}
	}
	return ParkingLevel{
		floor: floor,
		slots: slots,
	}
}

type ParkingSlot struct {
	column   int
	row      int
	occupied bool
}

func (p *ParkingSlot) IsFree() bool {
	return p.occupied
}

func (p *ParkingSlot) MarkOccupied() error {
	if p.occupied {
		p.occupied = false
	} else {
		return fmt.Errorf("slot already occupied")
	}
	return nil
}

func (p *ParkingSlot) MarkFree() {
	p.occupied = false
}

func NewParkingSlot(i, j int) *ParkingSlot {
	return &ParkingSlot{
		row:    i,
		column: j,
	}
}
