package main

import (
	"fmt"
	"log"
	"os"

	"github.com/IvanYaremko/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		fmt.Println("error")
	}

	appState := state{
		cfg: &configuration,
	}

	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("error provided less than two argumens")
		os.Exit(1)
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}

	if err := commands.run(&appState, cmd); err != nil {
		log.Fatal(err)
	}
}
