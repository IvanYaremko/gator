package main

import (
	"context"
	"errors"

	"github.com/IvanYaremko/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("provide <url> in args to unfollow")
	}
	url := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowsByIdParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	if err := s.db.DeleteFeedFollowsById(context.Background(), params); err != nil {
		return err
	}

	return nil
}
