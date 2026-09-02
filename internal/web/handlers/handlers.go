package handlers
import(
	"html/template"
	"net/http"
	"time"
	"techtrends.com/m/internal/web/templates"
)

var testtemplate = template.Must(template.ParseFS(templates.FS, "*.html"))

func Page(w http.ResponseWriter, r *http.Request){
	testtemplate.ExecuteTemplate(w, "index.html", nil)
}

func Fragment(w http.ResponseWriter, r *http.Request){
	testtemplate.ExecuteTemplate(w, "fragment.html", map[string]any{"Now": time.Now().Format("15:04:05")})
}