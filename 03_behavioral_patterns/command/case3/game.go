package main

type Game interface {
	start()
	end()
}

type GameEndCommand struct {
	game Game
}

func (c *GameEndCommand) execute() {
	c.game.end()
}
func (c *GameEndCommand) undo() {
	command := GameStartCommand{game: c.game}
	command.undo()
}

type GameStartCommand struct {
	game Game
}

func (c *GameStartCommand) execute() {
	c.game.start()
}
func (c *GameStartCommand) undo() {
	command := GameEndCommand{game: c.game}
	command.execute()
}
