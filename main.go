package main

import (
	"log"
	"os"

	"github.com/pedroaguia8/gator/internal/cli"
	"github.com/pedroaguia8/gator/internal/config"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Error: Not enough arguments provided")
	}

	cfg, err := config.Read()
	if err != nil {
		// TODO: instead create a new default config?
		log.Fatalf("Failed to read initial config: %v", err)
	}
	state := cli.State{}
	state.Config = &cfg

	commands := cli.Commands{Handlers: map[string]func(*cli.State, cli.Command) error{}}

	commands.Register("login", cli.HandlerLogin)

	command := cli.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.Run(&state, command)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

}
