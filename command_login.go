package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}

	userName := cmd.arguments[0]
	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("handler login couldn't set username: %w", err)
	}

	fmt.Println("✅", userName, "has been set!")

	return nil
}
