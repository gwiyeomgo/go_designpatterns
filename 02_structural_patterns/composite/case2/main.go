package main

type Component interface {
	getPrice() int
}
type Item struct {
	price int
}

func (i *Item) getPrice() int {
	return 0
	//fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type Bag struct {
	price string
}

func (b *Bag) getPrice() int {
	return 0
	//fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

func main() {

}
