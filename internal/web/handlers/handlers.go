package handlers
import(
	"html/template"
	"net/http"
	"techtrends.com/m/internal/web/templates"
)

var testtemplate = template.Must(template.ParseFS(templates.FS, "*.html"))

func Page(w http.ResponseWriter, r *http.Request){
	// testtemplate.ExecuteTemplate(w, "index.html", nil)
	testtemplate.ExecuteTemplate(w, "index.html", map[string]any{"Active": "home"})
}
func PastWeekStats(w http.ResponseWriter, r *http.Request){
	testtemplate.ExecuteTemplate(w, "pastweekstats.html", map[string]any{"Active": "pastweek"})
}