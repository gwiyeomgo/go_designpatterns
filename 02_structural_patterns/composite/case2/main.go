package main

import "fmt"

type Component interface {
	getPrice() int
}

func main() {
	lipstick := Item{name: "lipstick", price: 45000}
	sunglasses := Item{name: "sunglasses", price: 50000}

	bag := &Bag{}
	bag.add(&lipstick)
	bag.add(&sunglasses)

	fmt.Printf("가방안의 물건의 가격 합계는 %d", bag.getPrice())
}
