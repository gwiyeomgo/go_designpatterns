package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
퓨처 : 아직 알지 못하는 값에 대한 플레이스홀더를 제공

오랫동안 수행되는 블로킹 함수가 있는데 고루틴에서 이를 수행하고
채널과 함께 결과를 반환

두개의 행렬의 역곱을 개산하는 함수
*/

type Future interface {
	Result() (string, error)
}

type InnerFuture struct {
	once sync.Once
	wg   sync.WaitGroup

	res   string
	err   error
	resCh <-chan string
	errCh <-chan error
}

// sync.Once를 사용하여 보장됩니다. 또한, sync.WaitGroup를 사용하여 결과를 기다립니다.
func (f *InnerFuture) Result() (string, error) {
	f.once.Do(func() {

		f.wg.Add(1)
		defer f.wg.Done()

		f.res = <-f.resCh
		f.err = <-f.errCh
	})

	f.wg.Wait()

	return f.res, f.err
}

func SlowFunc(ctx context.Context) Future {
	resCh := make(chan string)
	errCh := make(chan error)

	go func() {
		//select 의 특징 모든 case가 계산
		timeout := time.After(2 * time.Second)
		select {
		case <-timeout:
			resCh <- " i slept for 2 seconds"
			errCh <- nil
		case <-ctx.Done():
			resCh <- ""
			errCh <- ctx.Err()
		}
	}()

	return &InnerFuture{resCh: resCh, errCh: errCh}
}

// const fmf를 사용해 타임아웃,데드라인을 적용할 수 있다
// https://www.whatap.io/ko/blog/171/
//https://hamait.tistory.com/748
/*
 비동기 작업을 수행하고 결과를 나중에 동기적으로 가져올 수 있는 패턴
결과를 필요한 시점에 가져올 수 있도록 하는 것
*/
func main() {
	ctx := context.Background()

	// SlowFunc 함수가 비동기 작업을 수행하여 Future를 얻습니다.
	//백그라운드에서 비동기 작업을 수행
	future := SlowFunc(ctx)

	res, err := future.Result()
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	fmt.Println(res)
}
