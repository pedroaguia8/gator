package cli

import "fmt"

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a username as parameter: login <username>")
	}

	username := cmd.Args[0]

	err := s.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set username: %w", err)
	}

	fmt.Printf("User %s has been set\n", username)
	return nil
}
