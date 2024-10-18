package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("provide <url> argument")
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("get feed error: %w", err)
	}

	feedFollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	resp, err := s.db.CreateFeedFollow(context.Background(), feedFollow)
	if err != nil {
		return fmt.Errorf("failed to creat follow: %w", err)
	}

	fmt.Println("Feed followed!")
	fmt.Println("Feed name:", resp.FeedName)
	fmt.Println("Followed by:", resp.UserName)

	return nil
}
