package main

import (
	"context"
	"fmt"

	"github.com/IvanYaremko/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
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
