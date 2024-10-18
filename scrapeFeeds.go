package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) error {

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

	postArg := database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    nextFeed.ID,
		Title:     feedPtr.Channel.Title,
		Description: sql.NullString{
			String: feedPtr.Channel.Description,
			Valid:  true,
		},
		Url: feedPtr.Channel.Link,
		PublishedAt: sql.NullTime{
			Time:  nextFeed.CreatedAt,
			Valid: true,
		},
	}

	if _, err := s.db.CreatePost(context.Background(), postArg); err != nil {
		return err
	}

	return nil
}
