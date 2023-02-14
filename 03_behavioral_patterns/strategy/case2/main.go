package main

func main() {
	missile := &MissileStrategy{}
	robot := initRobot("아톰", missile)
	robot.attack()
	punchStrategy := &PunchStrategy{}
	robot.setAttackStrategy(punchStrategy)
	robot.attack()
}
