package MetroCard

import (
	"sort"
)

type Log struct {
	journeys []PassengerJourney
}

func NewLog() *Log {
	return &Log{
		journeys: make([]PassengerJourney, 0),
	}
}

func (l *Log) AddJourney(journey *PassengerJourney) {
	l.journeys = append(l.journeys, *journey)
}

func (l *Log) IsReturnJourney(journey *PassengerJourney) bool {
	var count int
	for _, j := range l.journeys {
		if journey.Card.Equals(j.Card) {
			if j.Start.Equals(journey.End) {
				count++
			}
			if j.Start.Equals(journey.Start) {
				count--
			}
		}
	}
	return count%2 == 1
}

func (l *Log) GetCollectionAndDiscountOfStation(station Station) (float64, float64) {
	var collection float64
	var discount float64
	for _, j := range l.journeys {
		if j.Start.Equals(station) {
			collection += j.Collection
			discount += j.Discount
		}
	}
	return collection, discount
}

type JourneyPassengerCount struct {
	PassengerType PassengerType
	Count         int
}

func (l *Log) GetPassengerSummaryForStation(station Station) []JourneyPassengerCount {
	m := make(map[PassengerType]int)
	for _, j := range l.journeys {
		if j.Start.Equals(station) {
			if _, ok := m[j.Passenger.GetType()]; !ok {
				m[j.Passenger.GetType()] = 1
			} else {
				m[j.Passenger.GetType()] += 1
			}
		}
	}
	passengerCount := make([]JourneyPassengerCount, 0)
	for k, v := range m {
		passengerCount = append(passengerCount, JourneyPassengerCount{PassengerType: k, Count: v})
	}
	sort.Slice(passengerCount, func(i, j int) bool {
		if passengerCount[i].Count == passengerCount[j].Count {
			return passengerCount[i].PassengerType < passengerCount[j].PassengerType
		}
		return passengerCount[i].Count > passengerCount[j].Count
	})
	return passengerCount
}
