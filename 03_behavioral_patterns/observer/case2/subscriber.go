package main

// 옵저버
type Subscriber interface {
	handleMessage(message string)
}
