package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {

	if err := s.db.DeleteFeeds(context.Background()); err != nil {
		return fmt.Errorf("failed deleting feeds table: %w", err)
	}

	if err := s.db.DeteleUsers(context.Background()); err != nil {
		return fmt.Errorf("failed deleting users table: %w", err)
	}

	fmt.Println("removed all users from table")
	return nil
}
