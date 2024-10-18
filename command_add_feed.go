package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return errors.New("command requires <name> and <url> argument")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error get user in add feed handler: %w", err)
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

	fmt.Println("new feed added:", result)

	return nil
}
