package main

import "time"

type TourPlanBuilder struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
}

func newTourPlanBuilder() *TourPlanBuilder {
	return &TourPlanBuilder{}
}

func (b *TourPlanBuilder) setTitle(title string) {
	b.title = title
}

func (b *TourPlanBuilder) setNightsAndDays(nights int, days int) {
	b.nights = nights
	b.days = days
}

func (b *TourPlanBuilder) setStartDate(date time.Time) {
	b.startDate = date
}
func (b *TourPlanBuilder) setWhereToStay(stay string) {
	b.whereToStay = stay
}
func (b *TourPlanBuilder) getPlan() Plan {
	return Plan{
		title:       b.title,
		nights:      b.nights,
		whereToStay: b.whereToStay,
		startDate:   b.startDate,
		days:        b.days,
	}
}
