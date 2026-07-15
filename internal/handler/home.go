package handler

import (
	"log/slog"
	"net/http"

	"github.com/flintcraftstudio/nonprofit/internal/view"
)

// Home handles GET / and renders the home page. It is registered on the
// exact-match pattern "GET /{$}", so unknown paths fall through to NotFound
// rather than being served the homepage with a 200 (a soft 404).
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := view.HomePage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "err", err)
		}
	}
}

// NotFound is the catch-all for unmatched GET paths. It renders the gentle 404
// page with a real 404 status so search engines don't index dead links as
// duplicate homepages.
func NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		if err := view.NotFoundPage().Render(r.Context(), w); err != nil {
			slog.Error("render error", "page", "404", "err", err)
		}
	}
}
