package main

import (
	"fmt"
	"strings"
)

// 구상 데코레이터들
type SpamFiltering struct {
	comment IComment
}

func (p *SpamFiltering) addComment(comment string) {
	fmt.Println("SpamFiltering-----------------")
	//comment 에  http 이 포함한다면 추가 안함
	isSpam := strings.Contains(comment, "http")
	if !isSpam {
		p.comment.addComment(comment)
	}
}
