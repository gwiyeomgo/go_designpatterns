package main

import (
	"fmt"
	"sync"
	"time"
)

/*
팬인 : 다수의 입력 채널을 하나의 출력 채널로
*/

func FanIn(sources ...<-chan int) <-chan int {
	dest := make(chan int) // 공유 출력 채널 선언

	// 모든 sources 의 채널이 닫혔을 때 출력 채널을 자동으로 닫기 위해 사용
	var wg sync.WaitGroup

	wg.Add(len(sources)) //WaitGroup 은 입력 채널의 개수만큼 생성

	for _, ch := range sources { //각 입력 채널에 대해 고루틴을 시작
		go func(c <-chan int) {
			defer wg.Done() // 채널이 닫히면 waitGroup 으로 알려준다.

			for n := range c {
				dest <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()   // 모든 입력 채널이 닫힌 후
		close(dest) //출력 채널을 닫기 위한 고루틴 시작
	}()
	//단일 출력 채널 dest
	return dest
}
func main() {

	sources := make([]<-chan int, 0) //빈 채널 슬라이스를 생성

	for i := 0; i < 3; i++ {
		ch := make(chan int)
		sources = append(sources, ch) // 생성된 채널 추가

		go func() {
			defer close(ch) //고루틴 실행 끝나면 채널 닫기

			for i := 1; i <= 5; i++ {
				ch <- i
				time.Sleep(time.Second)
			}
		}()
	}

	dest := FanIn(sources...)
	for d := range dest {
		fmt.Print(d)
	}

}
