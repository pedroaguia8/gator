package cli

import (
	"context"
	"fmt"

	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a url as argument: unfollow <url>")
	}

	url := cmd.Args[0]
	feed, err := s.Db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("didn't find feed with given url: %w", err)
	}

	err = s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete the feed following: %w", err)
	}

	fmt.Printf("Not following feed %s as user %s anymore\n", feed.Name, user.Name)
	return nil
}
