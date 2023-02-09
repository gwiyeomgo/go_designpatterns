package main

type PostIterator struct {
	index int
	posts []*Post
}

func (u *PostIterator) hasPre() bool {
	if u.index > 0 {
		return true
	}
	return false
}

func (u *PostIterator) getPre() *Post {
	if u.hasPre() {
		post := u.posts[u.index-1]
		u.index--
		return post
	}
	return nil
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
