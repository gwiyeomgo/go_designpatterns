package main

import "fmt"

type Command interface {
	execute()
	undo()
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
