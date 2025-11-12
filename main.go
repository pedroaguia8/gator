package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pedroaguia8/gator/internal/database"
)
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

	// TODO: only allow register if user not set?
	cfg, err := config.Read()
	if err != nil {
		// TODO: instead create a new default config?
		log.Fatalf("Failed to read initial config: %v", err)
	}
	state := cli.State{}
	state.Cfg = &cfg
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	dbQueries := database.New(db)
	state.Db = dbQueries

	commands := cli.Commands{
		Handlers: map[string]func(*cli.State, cli.Command) error{},
	}

	commands.Register("login", cli.HandlerLogin)
	commands.Register("register", cli.HandlerRegister)
	commands.Register("reset", cli.HandlerReset)
	commands.Register("agg", cli.HandlerAgg)
	commands.Register("addfeed", cli.MiddlewareLoggedIn(cli.HandlerAddFeed))
	commands.Register("feeds", cli.HandlerFeeds)
	commands.Register("follow", cli.MiddlewareLoggedIn(cli.HandlerFollow))
	commands.Register("unfollow", cli.MiddlewareLoggedIn(cli.HandlerUnfollow))
	commands.Register("following", cli.MiddlewareLoggedIn(cli.HandlerFollowing))
	commands.Register("users", cli.HandlerUsers)
	commands.Register("browse", cli.MiddlewareLoggedIn(cli.HandlerBrowse))

	command := cli.Command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.Run(&state, command)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
