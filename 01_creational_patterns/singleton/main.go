package main

import (
	"designpatterns/01_creational_patterns/singleton/first"
	"fmt"
)

func main() {

	/*

		data := first.Settings1{}.GetInstance()
		data2 := first.Settings1{}.GetInstance()
		if data == data2 {
			fmt.Println("같음")
		} else {
			fmt.Println("v 다름")
		}

		data := first.Settings2{}.GetInstance()
		data2 := first.Settings2{}.GetInstance()
		if data == data2 {
			fmt.Println("v 같음")
		} else {
			fmt.Println("다름")
		}
	*/
	data := first.SyncSettings{}.GetInstance()
	data2 := first.SyncSettings{}.GetInstance()
	if data == data2 {
		fmt.Println("v 같음")
	} else {
		fmt.Println("다름")
	}
}

//https://refactoring.guru/design-patterns/singleton/go/example#example-1--syncOnce-go
