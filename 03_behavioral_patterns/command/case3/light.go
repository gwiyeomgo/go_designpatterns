package main

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
