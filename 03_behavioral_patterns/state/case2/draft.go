package main

import "errors"

//콘크리트 스테이트1

type Draft struct {
	onlineCourse *OnlineCourse //context
}

func (d Draft) addReview(review string, student *Student) error {
	return errors.New("드래프트 상태에서는 리뷰를 남길 수 없습니다.")
}
func (d Draft) addStudent(student Student) {
	students := d.onlineCourse.getStudents()
	if len(students) > 2 {

	}
}
