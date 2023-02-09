package main

type Iterator interface {
	hasNext() bool
	getNext() *Post
	hasPre() bool
	getPre() *Post
}
