package main

type Bag struct {
	components []Component
}

func (b *Bag) getPrice() int {
	sum := 0
	for _, composite := range b.components {
		sum += composite.getPrice()
	}
	return sum
}
func (b *Bag) add(c Component) {
	b.components = append(b.components, c)
}
