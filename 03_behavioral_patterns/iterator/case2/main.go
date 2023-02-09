package main

import (
	"fmt"
)

func main() {

	post1 := &Post{
		content: "첫번째 작성 글",
	}
	post2 := &Post{
		content: "두번째 작성 글",
	}
	post3 := &Post{
		content: "최근에 작성한 글",
	}
	board := &Board{
		posts: []*Post{post1, post2, post3},
	}
	//들어간 순서대로 순회
	iterator := board.getDefaultIterator()

	for iterator.hasNext() {
		post := iterator.getNext()
		fmt.Printf("post: %+v\n", post.content)
	}
	//가장 최신 글을 먼저 출력  ?

}
