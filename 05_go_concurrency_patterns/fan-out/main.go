package main

import (
	"fmt"
	"sync"
	"time"
)

/*
팬아웃 :입력 채널 하나에 수신된 메시지를 다수 출력 채널로 균등라게 분배

해당 채널이 닫힐 때까지 동일한 채널에서 여러 기능을 읽을 수 있습니다.

*/

// 이력 채널로부터 메세지를 수신하고
// 그것을 출력 채널에 공평하게 분배

func Split(source <-chan int, n int) []<-chan int {
	dests := make([]<-chan int, 0) //목적지 채널 슬라이스를 생성

	for i := 0; i < n; i++ {
		ch := make(chan int)
		dests = append(dests, ch)

		go func() {
			defer close(ch)

			for val := range source {
				ch <- val
			}
		}()

	}
	return dests
}
func main2() {
	//https://woojinger.tistory.com/82
	c := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			for n := range c {
				time.Sleep(1)
				fmt.Println(i, n)
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
func main() {
	source := make(chan int)
	dests := Split(source, 5) //5개 출력 채널을 갖음
	//1~10 까지의 값을 surce 전달 동시에 5개 dest로 값을 수신
	go func() {
		for i := 0; i <= 10; i++ {
			source <- i
		}
		close(source)
	}()

	var wg sync.WaitGroup
	wg.Add(len(dests)) // WaitGroup으로 모든 출력 채널이 닫힐 때까지 기다린다

	for i, ch := range dests {
		go func(i int, d <-chan int) {
			defer wg.Done() //wg.Wait() 가 잡고 있던 락을 해제 함 , 함수 종료

			for val := range d {
				fmt.Printf("#%d got %d\n", i, val)
			}
		}(i, ch)
	}
	wg.Wait()
}
