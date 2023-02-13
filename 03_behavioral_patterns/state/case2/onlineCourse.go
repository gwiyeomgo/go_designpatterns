package main

type OnlineCourse struct {
	state    State
	students []Student
	review   []string
}

func (o OnlineCourse) addStudent(student *Student) {
	o.state.addStudent(*student)
}

func (o OnlineCourse) addReview(review string, student *Student) {

}

func (o OnlineCourse) getState() State {
	return o.state
}

func (o OnlineCourse) getStudents() []Student {
	return o.students
}

func (o OnlineCourse) getReviews() []string {
	return o.review
}

func (o OnlineCourse) changeState(state State) {
	o.state = state
}
