package main

import "fmt"

type Command struct {
	name string
	args []string
}

type Commands struct {
	cmds map[string]func(*State, Command) error
}

func (c *Commands) run(s *State, cmd Command) error {
	callback, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("Command not found")
	}
	callback(s, cmd)
	return nil
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.cmds[name] = f
}
