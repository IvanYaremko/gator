package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return errors.New("command requires <name> and <url> argument")
	}

	name := cmd.arguments[0]
	url := cmd.arguments[1]

	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	result, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("error db create feed: %w", err)
	}

	newArgs := command{
		name:      "follow",
		arguments: []string{result.Url},
	}

	if err := handlerFollow(s, newArgs, user); err != nil {
		fmt.Println("failed to add auto follow")
	}

	fmt.Println("new feed added:", result)

	return nil
}
