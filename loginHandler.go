package main

import (
	"fmt"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login handler expects a single argument")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("User %v has been set properly", cmd.args[0])
	return nil
}
