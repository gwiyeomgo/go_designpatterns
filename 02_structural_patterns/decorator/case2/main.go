package main

// 클라이언트
type Client struct {
	comment IComment
}

func (c *Client) writeComment(comment string) {
	c.comment.addComment(comment)
}

func main() {

	comment := &Comment{}
	spamFilteringComment := &SpamFiltering{
		comment: comment,
	}

	spamFilteringAndTrimmingComment := &TrimmingComment{
		comment: spamFilteringComment,
	}

	client := Client{comment: spamFilteringAndTrimmingComment}
	client.writeComment("새해 복 많이 받아요!!")
	client.writeComment("아래 링크를 클릭하면 선물이...")
	client.writeComment("http://test.me")
}
