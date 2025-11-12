package cli

import (
	"context"
	"fmt"

	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerFollowing(s *State, _ Command, user database.User) error {
	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("could't retrieve current user's feed followings: %w", err)
	}

	fmt.Println("You are following these feeds:")
	for _, follow := range feedFollows {
		fmt.Println("* " + follow.FeedName)
	}
	return nil
}
