package main

import (
	"fmt"
)

type Repository struct {
	Name   string
	User   string
	Issues []Issue
}

func (r *Repository) print() {
	for _, i := range r.Issues {
		url := fmt.Sprintf("https://github.com/%s/%s/issues/%d", r.User, r.Name, i.Id)
		fmt.Println(url)
	}
}

type Issue struct {
	Id    int
	Title string
}

func (r *Issue) clone() *Issue {
	return &Issue{
		Id:    r.Id,
		Title: r.Title,
	}
}

func main() {
	// 스터디 프로젝트의 1개 이슈가 있다
	issue := Issue{
		Id:    1,
		Title: "개발 스프린트 1",
	}
	repository1 := &Repository{
		Name:   "스터디 프로젝트",
		User:   "gwiyeomgo",
		Issues: []Issue{issue},
	}

	repository1.print()
	repository2 := repository1

	// Modifying
	repository2.Name = "두번째 스터디 프로젝트"
	repository2.User = "gwiyeomgo"
	repository2.Issues = repository1.Issues
	repository2.print()
}
