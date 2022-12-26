package code

import "sync"

type SyncSettings struct {
}

var (
	once      sync.Once
	instance2 *SyncSettings // 포인터의 0 값이 nil
)

// Lazy Initialization: :게으른 초기화
func (s SyncSettings) LazyInit() *SyncSettings {
	// 아래 코드처럼 사용하면 race condition 으로 부터 보호됨
	once.Do(func() {
		instance2 = &s //thread safe
	})

	return instance2
}
