package main

// 여러 옵저버틀을 등록하고 조정
// 서브젝트의 상태가 변경되면 자신에게 증록되어
// 있는 옵저보를 순화하며 옵저버가 제공하는 특정한 메서드를 호출
type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
