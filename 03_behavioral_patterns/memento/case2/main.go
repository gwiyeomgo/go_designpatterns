package main

import "fmt"

// 클라이언트를 caretaker 라고 가정하고
func main() {
	game := Game{}
	game.setBlueTeamScore(10)
	game.setRedTeamScore(20)
	//순간 정보를 저장 스냅샷
	save := game.save()

	game.setBlueTeamScore(12)
	game.setRedTeamScore(22)

	game.restore(save)

	fmt.Println(game.getBlueTeamScore())
	fmt.Println(game.getRedTeamScore())
}
