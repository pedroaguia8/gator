package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, _ Command) error {
	err := s.Db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete all users: %w", err)
	}
	fmt.Println("All user records deleted")
	return nil
}
