package main

import "fmt"

type Student struct {
	name          string
	onlineCourses set
}

func (s Student) addPrivate(course *OnlineCourse) {
	s.onlineCourses.Add(course)
}
func (s Student) isAvailable(course *OnlineCourse) bool {
	return s.onlineCourses.Contain(course)
}

func (s Student) toString() string {
	return fmt.Sprintf("Student{ name= %s\n}", s.name)
}
