package main

import "fmt"

type Command interface {
	execute()
	undo()
}

type Light interface {
	on()
	off()
}

type LightOffCommand struct {
	light Light
}

func (c *LightOffCommand) execute() {
	c.light.off()
}
func (c *LightOffCommand) undo() {
	command := LightOnCommand{light: c.light}
	command.undo()
}

type LightOnCommand struct {
	light Light
}

func (c *LightOnCommand) execute() {
	c.light.on()
}
func (c *LightOnCommand) undo() {
	command := LightOffCommand{light: c.light}
	command.execute()
}

//game

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

type Button struct {
	commands Stack
}

func (b *Button) press(c Command) {
	c.execute()
	b.commands.push(c)
}

func (b *Button) undo() {
	if !b.commands.isEmpty() {
		c := b.commands.pop()
		c.undo()
	}
}

type Stack []Command

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
func (s *Stack) push(data Command) {
	*s = append(*s, data)
	fmt.Printf("%d pushed to stack\n", data)
}
func (s *Stack) pop() Command {
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return nil
	} else {
		top := len(*s) - 1
		data := (*s)[top]
		*s = (*s)[:top]
		return data
	}
}

func main() {

	var button Button
	button.press(&LightOnCommand{light: &IndoorLighting{}})
	button.press(&GameStartCommand{game: &PokemonGO{}})

	button.undo()
	button.undo()

}
