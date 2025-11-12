package rss

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pedroaguia8/gator/internal/database"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type DBLike interface {
	GetNextFeedToFetch(ctx context.Context) (database.Feed, error)
	MarkFeedFetched(ctx context.Context, arg database.MarkFeedFetchedParams) error
	CreatePost(ctx context.Context, arg database.CreatePostParams) (database.Post, error)
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	feed := RSSFeed{}
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, err
	}

	feed.decodeEscapedHtml()

	return &feed, nil
}

func (feed *RSSFeed) decodeEscapedHtml() {
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := 0; i < len(feed.Channel.Item); i++ {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
}

func ScrapeFeeds(ctx context.Context, db DBLike) error {
	feedToFetch, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("couldn't retrieve next feed to fetch: %w", err)
	}

	err = db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        feedToFetch.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't mark feed as fetched: %w", err)
	}

	feed, err := FetchFeed(ctx, feedToFetch.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	err = saveFeed(feed, feedToFetch.ID, ctx, db)
	if err != nil {
		return fmt.Errorf("couldn't save feed: %w", err)
	}

	return nil
}

func saveFeed(feed *RSSFeed, feedId uuid.UUID, ctx context.Context, db DBLike) error {
	for _, item := range feed.Channel.Item {
		parsedPubDate, err := parsePubDate(item.PubDate)
		if err != nil {
			log.Printf("couldn't parse 'published at' date %s for item %s of feed %s\n",
				item.PubDate, item.Title, feed.Channel.Title)
		}

		_, err = db.CreatePost(ctx, database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     item.Title,
			Url:       item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: parsedPubDate,
			FeedID:      feedId,
		})
		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url_key\"" {
				continue
			}
			log.Printf("couldn't save post: %s", err)
		}
	}
	return nil
}

func parsePubDate(pubDate string) (sql.NullTime, error) {
	parsedPubDate := sql.NullTime{}

	parsedTime, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		parsedPubDate.Valid = false
		return parsedPubDate, fmt.Errorf("couldn't parse date according to RFC1123Z layout: %w", err)
	}

	parsedPubDate.Valid = true
	parsedPubDate.Time = parsedTime
	return parsedPubDate, nil
}
