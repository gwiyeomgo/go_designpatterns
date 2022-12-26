package main

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
	//쓰레드세이프하지 않음
	/*
		data := first.SyncSettings{}.LazyInit()
		data2 := first.SyncSettings{}.LazyInit()
		if data == data2 {
			fmt.Println("v 같음")
		} else {
			fmt.Println("다름")
		}*/
	//

	/*//게으른 초기화
	go first.SyncSettings{}.LazyInit()
	go first.SyncSettings{}.LazyInit()
	*/
	/*
		go first.DoubleCheckedLockingSettings{}.LazyInit()
		go first.DoubleCheckedLockingSettings{}.LazyInit()
	*/
}

//https://refactoring.guru/design-patterns/singleton/go/example#example-1--syncOnce-go
