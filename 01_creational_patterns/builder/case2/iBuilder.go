package main

import "time"

// https://refactoring.guru/ko/design-patterns/builder/go/example
// 빌더 인터페이스
// 제품
type Plan struct {
	title       string
	nights      int
	days        int
	startDate   time.Time
	whereToStay string
}

type IBuilder interface {
	setTitle(string string)
	setWhereToStay(string string)
	setNightsAndDays(nights int, days int)
	setStartDate(date time.Time)
	getPlan() Plan
}

func getBuilder(builderType string) IBuilder {
	if builderType == "tour" {
		return newTourPlanBuilder()
	}

	if builderType == "travel" { //장기 여향
		return newTravelPlanBuilder()
	}
	return nil
}
