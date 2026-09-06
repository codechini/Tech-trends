package handlers
import(
	"html/template"
	"net/http"
	"techtrends.com/m/internal/web/templates"
	"fmt"
	"encoding/json"
)

type Post struct{
	ID int `json:"id"`
	Title string `json:"title"`
}

var testtemplate = template.Must(template.ParseFS(templates.FS, "*.html"))

func NewsFeed(w http.ResponseWriter, r *http.Request){
	resp,err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch news feed: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	
	var posts[]Post
	if err:= json.NewDecoder(resp.Body).Decode(&posts); err != nil{
		http.Error(w, fmt.Sprintf("Failed to decode news feed: %v", err), http.StatusInternalServerError)
		return
	}
	
	testtemplate.ExecuteTemplate(w, "newsfeed.html", map[string]any{"Active": "newsfeed", "NewsItems": posts})
}

func Page(w http.ResponseWriter, r *http.Request){
	// testtemplate.ExecuteTemplate(w, "index.html", nil)
	testtemplate.ExecuteTemplate(w, "index.html", map[string]any{"Active": "home"})
}
func PastWeekStats(w http.ResponseWriter, r *http.Request){
	testtemplate.ExecuteTemplate(w, "pastweekstats.html", map[string]any{"Active": "pastweek"})
}