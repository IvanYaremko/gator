package main

import (
	"context"
	"testing"
)

func TestFetchFeed(t *testing.T) {
	feed, err := fetchFeed(context.Background(), "https://blog.boot.dev/index.xml")
	if err != nil {
		t.Fatalf("fetchFeed returned an error: %v", err)
	}

	if feed.Channel.Title == "" {
		t.Error("expected non empty feed channel title, got an empty string")
	}
}
