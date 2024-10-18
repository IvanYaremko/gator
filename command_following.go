package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	data, err := s.db.GetFeedFollowsForUsers(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, item := range data {
		fmt.Println("Feeds following:")
		fmt.Println("feed name: ", item.FeedName)
	}

	return nil
}
