package main

import "time"

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) sapporoTrip() Plan {
	d.builder.setTitle(" [3박4일] 일본 삿포로 여행")
	d.builder.setStartDate(time.Date(2020, 10, 24, 12, 0, 0, 0, time.UTC))
	d.builder.setNightsAndDays(3, 4)
	d.builder.setWhereToStay("삿포로 A 호텔")
	return d.builder.getPlan()
}

func (d *Director) longBeachTrip() Plan {
	d.builder.setTitle("롱비치")
	d.builder.setStartDate(time.Date(2023, 11, 24, 12, 0, 0, 0, time.UTC))
	d.builder.setNightsAndDays(10, 11)
	d.builder.setWhereToStay("H 호텔")
	return d.builder.getPlan()
}
