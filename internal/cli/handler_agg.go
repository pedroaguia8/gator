package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pedroaguia8/gator/internal/rss"
)

func HandlerAgg(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("this command takes a duration as argument: agg <duration>. Ex.: agg 1h")
	}

	timeBetweenReqs := cmd.Args[0]
	durationBetweenRes, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return fmt.Errorf("couldn't parse time between requests string to duration: %w", err)
	}
	fmt.Printf("Collecting feeds every %s\n", durationBetweenRes.String())
	ticker := time.NewTicker(durationBetweenRes)
	for ; ; <-ticker.C {
		err := rss.ScrapeFeeds(context.Background(), s.Db)
		if err != nil {
			log.Printf("Error scraping feed: %v", err)
		}
	}
}
