package main

// context
// 컨텍스트에서 사용할 알고리즘을 클라이언트에서 선택한다
type Robot struct {
	name           string
	attackStrategy AttackStrategy
}

func initRobot(name string, a AttackStrategy) *Robot {
	return &Robot{
		name:           name,
		attackStrategy: a,
	}
}

func (r *Robot) setAttackStrategy(a AttackStrategy) {
	r.attackStrategy = a
}

func (r *Robot) getName() string {
	return r.name
}

func (r *Robot) attack() {
	r.attackStrategy.attack()
}
