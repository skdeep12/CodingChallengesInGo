package payment

import "context"

type Schedule interface {
	IsPayDay(context.Context) bool
}

type WeeklySchedule struct {
}

func NewWeeklySchedule() Schedule {
	return &WeeklySchedule{}
}

func (w *WeeklySchedule) IsPayDay(context.Context) bool {
	// check if current day is sunday
	return true
}

type BiWeeklySchedule struct {
}

func NewBiWeeklySchedule() Schedule {
	return &BiWeeklySchedule{}
}

type MonthlySchedule struct {
}

func NewMonthlySchedule() Schedule {
	return &MonthlySchedule{}
}

func (w *BiWeeklySchedule) IsPayDay(context.Context) bool {
	// check if current date is 15th or 31st
	return true
}

func (m *MonthlySchedule) IsPayDay(context.Context) bool {
	// check if current date is 31st
	return true
}
