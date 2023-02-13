package main

// 이벤트가 발생시 특정 메서드 호출
// 이 메서드들 안에서 옵저버가 해야할 일들 진행
type Observer interface {
	update(string)
	getID() string
}
