package main

func main() {
	missile := &MissileStrategy{}
	atom := initRobot("아톰", missile)
	atom.attack()

	punchStrategy := &PunchStrategy{}
	atom.setAttackStrategy(punchStrategy)
	atom.attack()
}
