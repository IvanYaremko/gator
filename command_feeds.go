package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("feeds handler: %w", err)
	}

	for _, feed := range feeds {

		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to retrieve user at feeds command: %w", err)
		}

		fmt.Println("Feed name:", feed.Name)
		fmt.Println("Feed url:", feed.Url)
		fmt.Println("Feed created by:", user.Name)
	}

	return nil
}
