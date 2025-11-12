package cli

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/pedroaguia8/gator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	limit := 2
	if len(cmd.Args) > 0 {
		limitInput, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: expected a number, but got '%s': %w", cmd.Args[0], err)
		}
		limit = limitInput
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Latest posts:\n\n")
	for _, post := range posts {
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("Description: %s\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("Published at: %s\n", post.PublishedAt.Time.Format(time.RFC1123Z))
		}
		fmt.Println()
	}
	return nil
}
