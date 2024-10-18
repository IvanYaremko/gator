package main

import (
	"errors"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
)

func handlerAgg(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return errors.New("provide time argument <1s | 1m |1h>")
	}

	duration, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(duration)

	for range ticker.C {
		scrapeFeeds(s, cmd, user)
	}

	return nil
}
