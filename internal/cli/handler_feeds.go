package cli

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, _ Command) error {
	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve feeds: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("%+v\n", feed)
	}
	return nil
}
