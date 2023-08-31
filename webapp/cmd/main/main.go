package main

import (
	"net/http"
	"os"

	"log/slog"

	"github.com/larschri/go-templates/webapp/internal/app"
	"github.com/larschri/go-templates/webapp/web"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	http.HandleFunc("/", app.MainHandler)
	http.HandleFunc("/error", app.ErrorHandler)
	http.HandleFunc("/time", app.TimeHandler)

	http.Handle("/static/", http.FileServer(http.FS(web.Static)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Info("exiting", "error", err)
	}
}
