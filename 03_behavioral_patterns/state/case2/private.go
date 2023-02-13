package main

// 콘크리트 스테이트2
type Private struct {
	onlineCourse *OnlineCourse
}

func (p Private) addItem(review string, student Student) {
}

func (p Private) addReview(review string, student Student) {
	students := p.onlineCourse.getStudents()
	reviews := p.onlineCourse.getReviews()
	if len(students) > 0 {
		for _, s := range students {
			if s == student {
				reviews = append(reviews, review)
			}
		}
	} else {
		//"프라이빗 코스를 수강하는 학생만 리뷰를 남길 수 있습니다."
	}
}

func (p Private) addStudent(student Student) {
	if student.isAvailable(p.onlineCourse) {
		p.onlineCourse.addStudent(&student)
	} else {

	}
}
