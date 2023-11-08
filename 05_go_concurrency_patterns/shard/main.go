package main

/*
샤딩 : 대규모 데이터 구조를 여러 개의 파티션으로 나눠 읽기/쓰기 락의 영향 범위 최소화

맵은 동시에 접근하는 경우 안전하지 않은 타입이기 때문에
어쩔 수 없이 락을 통한 동기화 구조를 사용해야함

여러 고루틴이 표준 맵에 동시에 액세스하고 수정하는 경우 문제 발생

각 샤드는 별도의 락을 가진 별도의 맵처럼 작동하며
고루틴은 서로 블로킹하지 않고 서로 다른 샤드를 동시에 액세스할 수 있습니다
*/
import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type Shard struct {
	sync.RWMutex                        //읽기 쓰기 모두에 대해서 락을 걸 수 있다
	m            map[string]interface{} //샤드의 데이터를 갖고 있음
}

type SharedMap []*Shard

func NewSharedMap(nshards int) SharedMap {
	shards := make([]*Shard, nshards) //슬라이스 초기화

	for i := 0; i < nshards; i++ {
		shard := make(map[string]interface{})
		shards[i] = &Shard{m: shard}
	}
	return shards
}

func (m SharedMap) getShardIndex(key string) int {
	checksum := sha1.Sum([]byte(key)) // crypto/sha1 패키지 sum 호출
	hash := int(checksum[17])         //임의 바이트를 고름
	return hash % len(m)              //
}

func (m SharedMap) getShard(key string) *Shard {
	index := m.getShardIndex(key)
	return m[index]
}

func (m SharedMap) Get(key string) interface{} {
	shard := m.getShard(key)
	shard.RLock()
	defer shard.RUnlock()

	return shard.m[key]
}

func (m SharedMap) Set(key string, value interface{}) {
	shard := m.getShard(key)
	shard.Lock()
	defer shard.Unlock()

	shard.m[key] = value
}
func (m SharedMap) Keys() []string {
	var keys []string
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard) {
			s.RLock()

			defer wg.Done()
			defer s.RUnlock()

			for key, _ := range s.m {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}
		}(shard)
	}

	wg.Wait()

	return keys
}

func main() {
	//다섯 개의 샤드를 가진 SharedMap를 만든다
	m := NewSharedMap(5)
	keys := []string{"alpha", "beta", "gamma"}

	for i, k := range keys {
		m.Set(k, i+1)

		fmt.Printf("%5s: shard=%d value=%d\n",
			k, m.getShardIndex(k), m.Get(k))
	}

	fmt.Println(m.Keys())
}
