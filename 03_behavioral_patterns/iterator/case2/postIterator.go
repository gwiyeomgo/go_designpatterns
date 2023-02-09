package main

type PostIterator struct {
	index int
	posts []*Post
}

func (u *PostIterator) hasNext() bool {
	if u.index < len(u.posts) {
		return true
	}
	return false

}
func (u *PostIterator) getNext() *Post {
	if u.hasNext() {
		post := u.posts[u.index]
		u.index++
		return post
	}
	return nil
}
