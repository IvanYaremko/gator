package main

import (
	"context"
	"fmt"
)

func hanlderUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("handler list: %w", err)
	}

	for _, user := range users {
		output := fmt.Sprintf("* %s", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			output = fmt.Sprintf("%s (current)", output)
		}
		fmt.Println(output)
	}

	return nil
}
