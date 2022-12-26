package first

type Settings2 struct {
}

var (
	instance *Settings2
)

func (s Settings2) GetInstance() *Settings2 {
	if instance == nil {
		instance = &s
	}
	return instance
}

//멀티쓰레드를 사용하는 경우 안전하지 않다
