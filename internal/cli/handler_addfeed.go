package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerAddFeed(s *State, cmd Command) error {
	currentUser := s.Cfg.CurrentUserName
	user, err := s.Db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error retrieving current user: %w", err)
	}

	if len(cmd.Args) < 2 {
		return fmt.Errorf("this command takes 2 arguments: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Printf("Feed created: %+v\n", feed)

	_, err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error following the new feed for the current suer: %w", err)
	}
	return nil
}
