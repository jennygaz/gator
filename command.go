package main

import "errors"

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	registeredCommands map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registeredCommands[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
