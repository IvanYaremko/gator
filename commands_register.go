package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/IvanYaremko/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return errors.New("handle register has 0 args")
	}

	name := cmd.arguments[0]

	if _, err := s.db.GetUser(context.Background(), name); err == nil {
		log.Fatal("User is already registered")
	}

	createParams := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      name,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	user, err := s.db.CreateUser(context.Background(), createParams)
	if err != nil {
		log.Fatal("Error creating user")
	}
	s.cfg.SetUser(user.Name)
	fmt.Println("user was created")
	fmt.Println("name: ", user.Name)
	fmt.Println("created at: ", user.CreatedAt)

	return nil
}
