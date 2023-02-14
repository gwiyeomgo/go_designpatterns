package main

type Robot struct {
	name           string
	attackStrategy IAttackStrategy
}

func initRobot(name string, a IAttackStrategy) *Robot {
	return &Robot{
		name:           name,
		attackStrategy: a,
	}
}

func (r *Robot) setAttackStrategy(a IAttackStrategy) {
	r.attackStrategy = a
}

func (r *Robot) getName() string {
	return r.name
}

func (r *Robot) attack() {
	r.attackStrategy.attack()
}
