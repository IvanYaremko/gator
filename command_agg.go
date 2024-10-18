package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	_, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("handler agg error: %w", err)
	}

	return nil
}
