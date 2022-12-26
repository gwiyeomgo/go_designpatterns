package first

type Settings1 struct {
}

func (s Settings1) GetInstance() *Settings1 {
	return &s
}
