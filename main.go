package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/IvanYaremko/gator/internal/config"
	"github.com/IvanYaremko/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	configuration, err := config.ReadConfig()
	if err != nil {
		fmt.Println("error")
	}

	db, err := sql.Open("postgres", configuration.DbURL)
	if err != nil {
		fmt.Println("error sql.Open")
	}
	dbQueries := database.New(db)

	appState := state{
		db:  dbQueries,
		cfg: &configuration,
	}

	commands := commands{
		cmds: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", hanlderUsers)
	commands.register("agg", handlerAgg)

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
