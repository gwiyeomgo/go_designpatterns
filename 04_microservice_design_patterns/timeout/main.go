package main

import (
	"context"
	"fmt"
	"time"
)

// https://www.vinsguru.com/timeout-pattern/
type WithContext func(context.Context, string) (string, error)
type SlowFunction func(string) (string, error)

// SlowFunction 을 클로저로 감싸고
func Timeout(f SlowFunction) WithContext {
	return func(ctx context.Context, arg string) (string, error) {
		ch := make(chan struct {
			result string
			err    error
		}, 1)

		go func() {
			res, err := f(arg)
			ch <- struct {
				result string
				err    error
			}{res, err}
		}()

		select {
		case res := <-ch:
			return res.result, res.err
		case <-ctx.Done():
			return "", ctx.Err()
		}
	}
}

func main() {
	slow := func(string2 string) (string, error) {
		time.Sleep(3 * time.Second)
		return "slow 호출", nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	timeout := Timeout(slow)
	res, err := timeout(ctx, "some input")
	if err != nil {
		fmt.Printf("에러: %v\n", err)
	} else {
		fmt.Printf("결과: %s\n", res)
	}
}
