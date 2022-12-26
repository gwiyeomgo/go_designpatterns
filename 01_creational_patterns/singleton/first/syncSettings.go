package first

import "sync"

type SyncSettings struct {
}

var (
	once      sync.Once
	instance2 *SyncSettings
)

// java synchronized 키워드 메서드에 한번에 한 쓰레드만 들어오도록 => sync pkg
// synchronized 단점은 getInstance 호출시 동기화 처리 작업으로 성능에 불이득
//이른 초기화 쓰레드 세이프...모르겠고 평소 사용하는 싱글톤 패턴

func (s SyncSettings) GetInstance() *SyncSettings {
	once.Do(func() {
		instance2 = &s
	})
	return instance2
}
