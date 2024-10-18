package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeteleUsers(context.Background())

	if err != nil {
		return fmt.Errorf("failed deleting users: %w", err)
	}

	fmt.Println("removed all users from table")
	return nil
}
