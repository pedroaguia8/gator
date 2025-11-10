package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a username as parameter: register <username>")
	}

	username := cmd.Args[0]

	user, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	})
	if err != nil {
		log.Fatalf("Registering new user failed: %v", err)
	}

	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set new user: %w", err)
	}

	fmt.Printf("New user %s created\n", username)
	log.Printf(
		"New user created:\n  ID: %v\n  Name: %v\n  CreatedAt: %v\n  UpdatedAt: %v\n",
		user.ID,
		user.Name,
		user.CreatedAt,
		user.UpdatedAt,
	)

	return nil
}
