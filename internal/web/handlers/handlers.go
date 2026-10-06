package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"techtrends.com/m/internal/ingester"
	"techtrends.com/m/internal/web/templates"
)

var testtemplate = template.Must(template.ParseFS(templates.FS, "*.html"))

// NewsFeed delegates RSS fetching/parsing to the ingester package and renders the result.
func NewsFeed(w http.ResponseWriter, r *http.Request) {
	posts, err := ingester.ParseRSSFeed(r.Context(), "https://www.theverge.com/rss/index.xml")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing RSS feed: %v", err), http.StatusBadGateway)
		return
	}
	testtemplate.ExecuteTemplate(w, "newsfeed.html", map[string]any{"Active": "newsfeed", "NewsItems": posts})
}

func Page(w http.ResponseWriter, r *http.Request) {
	// testtemplate.ExecuteTemplate(w, "index.html", nil)
	testtemplate.ExecuteTemplate(w, "index.html", map[string]any{"Active": "home"})
}
func PastWeekStats(w http.ResponseWriter, r *http.Request) {
	testtemplate.ExecuteTemplate(w, "pastweekstats.html", map[string]any{"Active": "pastweek"})
}
