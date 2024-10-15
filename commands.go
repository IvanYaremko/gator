package main

import (
	"fmt"
)

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	commandFunction, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("run command not found: %s", cmd.name)
	}

	if err := commandFunction(s, cmd); err != nil {
		return fmt.Errorf("run command %s: %w", cmd.name, err)
	}
	return nil
}
