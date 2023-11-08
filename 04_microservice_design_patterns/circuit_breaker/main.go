package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

// 클로저는 부모 함수의 변수에 액세스할 수 있는 중첩 함수
func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	//변수 선언과 초기화 구문은 프로그램이 실행될 때 한 번만 실행,프로그램이 시작될 때 한 번 초기화
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex //읽기-쓰기 뮤텍스
	fmt.Println("lastAttempt", lastAttempt)

	return func(ctx context.Context) (string, error) {
		m.RLock() // 읽기 잠금 수행
		d := consecutiveFailures - int(failureThreshold)
		if d >= 0 {
			// 현재 시각에 (2 << d)초를 추가하고 재시도를 언제 할지 결정
			//2 << d는 2의 d 제곱
			// 최초 호출된 시점 시간에 4초 를 더한 시간
			shouldRetryAt := lastAttempt.Add(time.Second * (2 << d))

			if !time.Now().After(shouldRetryAt) {
				m.RUnlock() // 읽기 잠금 풀기
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock() // 읽기 잠금 풀기
		response, err := circuit(ctx)

		m.Lock() //공유 자원 잠금
		defer m.Unlock()
		lastAttempt = time.Now()
		//서비스 호출이 실패
		if err != nil {
			consecutiveFailures++
			return response, err
		}
		//서비스 호출 성공시 실패 횟수 초기화
		consecutiveFailures = 0
		return response, nil
	}
}

func main() {
	// 예제 Circuit 함수
	exampleCircuit := func(ctx context.Context) (string, error) {
		// 실제 서비스 호출 또는 동작을 시뮬레이트하고 예제 응답을 반환
		return "", errors.New("exampleCircuit err")
	}

	// Circuit Breaker를 적용한 서비스 호출
	breakerCircuit := Breaker(exampleCircuit, 3) // 예: 3번 연속 실패하면 브레이크
	for i := 0; i < 5; i++ {
		response, err := breakerCircuit(context.Background())
		if err != nil {
			fmt.Printf("Service call failed: %v\n", err)
		} else {
			fmt.Printf("Service response: %s\n", response)
		}
		time.Sleep(1 * time.Second) // 일부 실패 시간 간격을 기다림
	}

}
