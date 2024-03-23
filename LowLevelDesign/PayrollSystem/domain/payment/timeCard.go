package payment

import "time"

type TimeCard struct {
	Hours     int
	CreatedAt time.Time
}

func NewTimeCard(hours int, recordDate time.Time) TimeCard {
	return TimeCard{
		Hours:     hours,
		CreatedAt: recordDate,
	}
}
