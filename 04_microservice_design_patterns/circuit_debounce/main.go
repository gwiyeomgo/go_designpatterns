package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 서킷 : 조절 기능 수행
// 디바운스: 서킷과 동일한 함수 시그니처를 가진 클로저

type Circuit func(ctx context.Context) (string, error)

// 펑션 퍼스트 : 최초 응답을 캐시하여 이후의 응답에 사용
func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer func() {
			threshold = time.Now().Add(d)
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			return result, err
		}

		result, err = circuit(ctx)

		return result, err
	}
}

// 펑션 라스트 : 일련의 요청에 대해 바로 내부 함수를 호출하지 않고 마지막 요청까지 기다린다
// 검색창에 입력 값을 넣는 동안 매번 함수를 호출하지 않고 입력이 잠시 멈췄을 때 입력된 값을 사용하여 함수 호출
// https://jake-seo-dev.tistory.com/651#Debounce%---%EB%--%--%EB%B-%--%EC%-A%B-%EC%-A%A--%--%EC%--%--%--Throttling%---%EC%--%B-%EB%A-%-C%ED%-B%--%EB%A-%---%--%EC%-D%B-%EB%-E%--%-F
func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time = time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var m sync.Mutex

	return func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()

		threshold = time.Now().Add(d)

		once.Do(func() {
			//ticker는 1초마다 시간을 확인하여 마지막 호출 시간(threshold)을 갱신할 역할을 합니다.
			ticker = time.NewTicker(time.Millisecond * 1000)
			fmt.Println("새로운 조회")
			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()

				for {
					select {
					case <-ticker.C:
						m.Lock()
						if time.Now().After(threshold) {

							result, err = circuit(ctx)
							m.Unlock()
							return
						}
						m.Unlock()
					case <-ctx.Done():
						//wrapped 함수를 통해 호출될 때,
						//threshold 를 갱신하고 마지막 호출 시간으로부터 d 이상 경과하지 않으면 대기합니다.
						//따라서 ctx.Done()이 출력되는 것은 디바운싱 시간 내에 새로운 호출이 발생하기 때문입니다.
						fmt.Println("시간내 조회")
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		})
		return result, err
	}
}

func DebounceLastEx() {
	exampleCircuit := func(ctx context.Context) (string, error) {
		return time.Now().String(), nil
	}
	//특정 시간 동안의 호출을 무시하고 마지막 호출의 결과를 반환
	//wrapped 함수는 디바운싱 시간(time.Second*1) 내에 다시 호출되면
	//이전 결과를 반환하고 새로운 호출은 무시됩니다.
	wrapped := DebounceLast(exampleCircuit, time.Millisecond*4000)
	//wrapped 함수를 5번 호출하고 1초마다 호출하는 동안에 ctx.Done이 출력
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		response, err := wrapped(context.Background())

		if err != nil {
			fmt.Printf("Service call failed: %v\n", err)
		} else {
			fmt.Printf("Service response: %s\n", response)
		}

		time.Sleep(3000 * time.Millisecond)
	}
}

// https://www.oreilly.com/library/view/cloud-native-go/9781098156411/ch04.html
func DebounceFirstEx() {
	// 예제 Circuit 함수 ->  2초안에 1번만 호출됨
	exampleCircuit := func(ctx context.Context) (string, error) {
		return time.Now().String(), nil
	}
	// 최초 응답을 캐시하고, 캐시가 존재하는 경우 threshold 시간 내에 호출되면 이전 응답을 반환
	//duration 값을 증가시키면 더 오랫동안 이전 응답이 캐싱
	//duration 값을 줄이면 더 빈번한 서비스 호출
	wrapped := DebounceFirst(exampleCircuit, time.Millisecond*4000)
	for i := 0; i < 5; i++ {
		response, err := wrapped(context.Background())
		if err != nil {
			fmt.Printf("Service call failed: %v\n", err)
		} else {
			fmt.Printf("Service response: %s\n", response)
		}
		time.Sleep(2000 * time.Millisecond)
	}
}

func main() {

	DebounceLastEx()
	//DebounceFirstEx()

}
