package cli

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a username as parameter: login <username>")
	}

	username := cmd.Args[0]

	_, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find registered user: %w", err)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set username: %w", err)
	}

	fmt.Printf("User %s has been set\n", username)
	return nil
}
