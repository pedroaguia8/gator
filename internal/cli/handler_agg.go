package cli

import (
	"context"
	"fmt"

	"github.com/pedroaguia8/gator/internal/rss"
)

func HandlerAgg(_ *State, _ Command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Println(feed)
	return nil
}
