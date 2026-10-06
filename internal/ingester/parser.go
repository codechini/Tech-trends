// Parse RSS feed from url parameter and return the parsed feed to nlp.go for processing.
package ingester

import (
	"context"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)
type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Link      string `json:"link"`
	Published string `json:"published"`
	Description string `json:"description"`
	GUID      string `json:"guid"`
}

func ParseRSSFeed(ctx context.Context, feedURL string) ([]Post, error) { //should take list of strings as input
	if feedURL == "" {
		return nil, fmt.Errorf("parse rss: empty feedURL")
	}
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(feedURL, ctx)
	if err != nil {
		return nil, fmt.Errorf("parse rss %q: %w", feedURL, err)
	}
	if feed == nil || len(feed.Items) == 0 {
		return []Post{}, nil
	}

	posts := make([]Post, 0, len(feed.Items))
	for i, item := range feed.Items {
		if item == nil {
			continue
		}
		posts = append(posts, Post{
			ID:        i + 1,
			Title:     item.Title,
			Link:      item.Link,
			Published: item.Published,
			Description: item.Description,
			GUID:      item.GUID,
		})
	}
	return posts, nil
}
// need to integrate with Cadence to keep track of the state of the function.
// next nee d to pass the parsed content to nlp.py toextract keywords.