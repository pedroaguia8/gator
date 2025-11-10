package cli

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, _ Command) error {
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve users: %w", err)
	}

	loggedInUser := s.Cfg.CurrentUserName

	for _, user := range users {
		if loggedInUser == user.Name {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
