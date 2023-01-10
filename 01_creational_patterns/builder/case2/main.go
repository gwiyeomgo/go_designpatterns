package main

import "fmt"

// 클라이언트
func main() {
	tourBuilder := getBuilder("tour")     //짧은여행
	travelBuilder := getBuilder("travel") //장

	director := newDirector(tourBuilder)
	tour := director.sapporoTrip()

	fmt.Printf("Normal Plan title: %s\n", tour.title)
	fmt.Printf("Normal House whereToStay: %s\n", tour.whereToStay)
	fmt.Printf("Normal House startDate: %s\n", tour.startDate)
	fmt.Printf("Normal House nights: %d\n", tour.nights)
	fmt.Printf("Normal House days: %d\n", tour.days)

	director.setBuilder(travelBuilder)
	travel := director.longBeachTrip()

	fmt.Printf("Normal Plan title: %s\n", travel.title)
	fmt.Printf("Normal House whereToStay: %s\n", travel.whereToStay)
	fmt.Printf("Normal House startDate: %s\n", travel.startDate)
	fmt.Printf("Normal House nights: %d\n", travel.nights)
	fmt.Printf("Normal House days: %d\n", travel.days)

}
