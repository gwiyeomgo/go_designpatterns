package main

func main() {
	onlineCourse := &OnlineCourse{}
	student1 := &Student{
		name: "testName",
	}
	student2 := &Student{
		name: "testName2",
	}
	student2.addPrivate(onlineCourse)

	onlineCourse.addStudent(student1)
	onlineCourse.changeState(&Private{onlineCourse: onlineCourse})

	onlineCourse.addReview("hello", student1)

	onlineCourse.addStudent(student2)
}
