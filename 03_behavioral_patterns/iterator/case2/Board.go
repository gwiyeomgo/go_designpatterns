package main

type Board struct {
	posts []*Post
}

func (p *Board) getDefaultIterator() Iterator {
	return &PostIterator{
		posts: p.posts,
	}
}
