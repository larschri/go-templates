package app

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/larschri/go-templates/webapp/internal/htmx"
	"github.com/larschri/go-templates/webapp/web"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	HandleTemplate(w, web.Templates.Lookup("main.html"), nil)
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", time.Now().Format(time.DateTime))
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	HandleError(w, "An error looks like this")
}

// HandleTemplate renders inner inside the outer.html
func HandleTemplate(w http.ResponseWriter, inner *template.Template, c any) {
	var buf bytes.Buffer
	if err := inner.Execute(&buf, c); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	web.Templates.Lookup("outer.html").Execute(w, template.HTML(buf.String()))
}

// HandleError renders error.html with an error message
func HandleError(w http.ResponseWriter, content any) {
	w.Header().Set(htmx.HXRetarget, "body")
	w.Header().Set(htmx.HXReswap, "afterbegin")
	web.Templates.Lookup("error.html").Execute(w, content)
}
