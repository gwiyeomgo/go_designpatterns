package main

import (
	"fmt"
	"time"
)

type GameService interface {
	startGame()
}

// realSubject
type DefaultGameService struct {
}

func (c *DefaultGameService) startGame() {
	fmt.Println("이 자리에 오신 여러분을 진심으로 환영합니다.")
}

// 프록시
type GameServiceProxy struct {
	gameService GameService
}

func (g *GameServiceProxy) startGame() {
	start := time.Now()
	if g.gameService == nil {
		g.gameService = &DefaultGameService{}
	}
	time.Sleep(8 * time.Second)
	g.gameService.startGame()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func main() {
	gameService := GameServiceProxy{}
	gameService.startGame()
}
