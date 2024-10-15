package main

import (
	"fmt"

	"github.com/IvanYaremko/gator/internal/config"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("error")
	}
	cfg.SetUser("ivan")

	cfg, err = config.ReadConfig()
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println(cfg)
}
