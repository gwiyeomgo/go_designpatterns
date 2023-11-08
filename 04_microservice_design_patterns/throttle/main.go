package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//특정 주기의 시간 동안 함수가 호출될 수 있는 횟수를 제한
/*
사용자는 초당 10회의 서비스 요청만 할 수 있다
고객은 특정 함수를 0.5초에 한 번만 호출할 수 있다
특정 계정에 대해 24시간 동안 3번까지만 로그인 실패를 허용
*/
/*
충분한 토큰을 사용할 수 있을 때 요청을 대기열에 넣기

*/
type Effector func(context.Context) (string, error)

/*
	refill =주어진 시간 간격(d 변수가 나타내는 시간 동안)에 다시 토큰을 채울 양

, max를 5로 설정하고 refill을 2로 설정하면,
초기에 5개의 토큰이 생성됩니다.
그런 다음 d로 지정된 시간 간격 동안마다 2개의 토큰이 생성
*/
//빈도 제한 로직=> 토큰 버킷 알고리즘 = 최대 토큰 수를 저자할 수 있는 버킷의 유사성 활용
func Throttle(e Effector, max uint, refill uint, d time.Duration) Effector {
	var tokens = max
	var once sync.Once
	//refill은 처음 d의 주기에만 적용
	return func(ctx context.Context) (string, error) {
		if ctx.Err() != nil {
			return "", ctx.Err()
		}

		once.Do(func() {
			ticker := time.NewTicker(d)

			go func() {
				defer ticker.Stop()

				for {
					select {
					case <-ctx.Done():
						return

					case <-ticker.C:

						t := tokens + refill
						//토큰 수를 최대 5개로 제한하면서 Effector 함수를 호출할 수 있습니다.
						if t > max {
							t = max
						}
						tokens = t
					}
				}
			}()
		})

		//토큰 수가 0보다 작으면 쓰로틀링에 의해 "too many calls" 오류가 반환됩니다
		if tokens <= 0 {
			//에버반환 : http 429 너무 많은 요청
			return "", fmt.Errorf("too many calls")
		}

		tokens--

		return e(ctx)
	}
}

func main() {
	// 예제 Effector 함수 생성
	exampleEffector := func(ctx context.Context) (string, error) {
		return "Effector 호출", nil
	}

	// Throttle 함수를 사용하여 호출 속도를 제한
	throttledEffector := Throttle(exampleEffector, 5, 2, time.Second)

	// 몇 번 호출해보기
	for i := 0; i < 10; i++ {
		result, err := throttledEffector(context.Background())
		if err != nil {
			fmt.Printf("에러: %v\n", err)
		} else {
			fmt.Printf("결과: %s\n", result)
		}
	}

}
