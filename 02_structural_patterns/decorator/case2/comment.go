package main

import "fmt"

// 컴포넌트
type IComment interface {
	addComment(comment string)
}

// 기초 데코레이터
type Comment struct {
	comment IComment
}

func (c *Comment) addComment(comment string) {
	fmt.Println("Comment-----------------------")
	fmt.Println(comment)
}
