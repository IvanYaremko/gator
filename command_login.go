package main

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("username is required")
	}

	userName := cmd.arguments[0]

	if _, err := s.db.GetUser(context.Background(), userName); err != nil {
		log.Fatal("user is not registered")
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("handler login couldn't set username: %w", err)
	}

	fmt.Println("✅", userName, "has been set!")

	return nil
}
