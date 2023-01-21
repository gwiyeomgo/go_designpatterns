package main

import (
	"fmt"
	"strings"
)

// 구상 데코레이터들
type TrimmingComment struct {
	comment IComment
}

func trim(comment string) string {
	//반복 횟수에 -1 = 일치하는 모든 문자를 지정한 치환 문자로 변경
	return strings.Replace(comment, "...", "", -1)
}
func (t *TrimmingComment) addComment(comment string) {
	fmt.Println("TrimmingComment---------------")
	t.comment.addComment(trim(comment))
}
