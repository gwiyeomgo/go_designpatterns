package main

import "fmt"

type Component interface {
	getPrice() int
}

// 어떤 새 종류의 컴포넌트가 생겨도 클라이언트 코드는 바뀌지 않는다
// 오픈 클로즈드 프린스펄
// 확장에는 열려있다 변경에는 닫혀있다.
func printPrice(component Component) {
	fmt.Println(component.getPrice())
}

// 전체나 부분이나 클라이언트 입장에서는 동일하게
func main() {
	lipstick := Item{name: "lipstick", price: 45000}
	sunglasses := Item{name: "sunglasses", price: 50000}

	bag := &Bag{}
	bag.add(&lipstick)
	bag.add(&sunglasses)

	printPrice(bag)
	//fmt.Printf("가방안의 물건의 가격 합계는 %d", bag.getPrice())
}
