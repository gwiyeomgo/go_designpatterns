package main

type Item struct {
	name  string
	price int
}

func (i *Item) getPrice() int {
	return i.price
}
