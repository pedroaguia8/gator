package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, _ Command) error {
	currentUser, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't retrieve current user: %w", err)
	}

	feedFollows, err := s.Db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("could't retrieve current user's feed followings: %w", err)
	}

	fmt.Println("You are following these feeds:")
	for _, follow := range feedFollows {
		fmt.Println("* " + follow.FeedName)
	}
	return nil
}
