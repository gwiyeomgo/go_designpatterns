package code

import "sync"

type DoubleCheckedLockingSettings struct {
}

var (
	lock      sync.Mutex
	instance3 *DoubleCheckedLockingSettings // 포인터의 0 값이 nil
)

func (d DoubleCheckedLockingSettings) LazyInit() *DoubleCheckedLockingSettings {
	//Double checked locking in golang
	if instance3 != nil { // 1st check
		return instance3
	}

	lock.Lock()
	defer lock.Unlock()

	if instance3 != nil {
		return instance3
	}
	instance3 = &d //thread safe
	return instance3
}
