package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerFollow(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a url as argument: follow <url>")
	}

	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get current user: %w", err)
	}

	url := cmd.Args[0]
	feed, err := s.Db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("didn't find feed with given url: %w", err)
	}

	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create the feed following: %w", err)
	}

	fmt.Printf("Following feed %s as user %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
