package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
)

func scrapeFeeds(s *state, cmd command, user database.User) error {

	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	arg := database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
	}
	if err := s.db.MarkFeedFetched(context.Background(), arg); err != nil {
		return err
	}

	feedPtr, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Println("scrapeFeeds ptr title", feedPtr.Channel.Title)
	return nil
}
