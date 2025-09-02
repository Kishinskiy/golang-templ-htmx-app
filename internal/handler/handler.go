package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)
type Dependencies struct {
	AssetsFS http.FileSystem
}

type hadlerFunc func(http.ResponseWriter, *http.Request) error

func RegisterRouters(r *chi.Mux, deps Dependencies) {
	home := homeHandler{}
	r.Get("/", handler(home.handleIndex))
	r.Get("/about", handler(home.handleAbout))

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(deps.AssetsFS)))
}


func handler(h hadlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("Error during request", slog.String("error", err.Error()))
}
