package main

import "time"

type TravelPlanBuilder struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
}

func newTravelPlanBuilder() *TravelPlanBuilder {
	return &TravelPlanBuilder{}
}

func (b *TravelPlanBuilder) setTitle(title string) {
	b.title = title
}

func (b *TravelPlanBuilder) setNightsAndDays(nights int, days int) {
	b.nights = nights
	b.days = days
}

func (b *TravelPlanBuilder) setStartDate(date time.Time) {
	b.startDate = date
}
func (b *TravelPlanBuilder) setWhereToStay(stay string) {
	b.whereToStay = stay
}
func (b *TravelPlanBuilder) getPlan() Plan {
	return Plan{
		title:       b.title,
		nights:      b.nights,
		whereToStay: b.whereToStay,
		startDate:   b.startDate,
		days:        b.days,
	}
}
