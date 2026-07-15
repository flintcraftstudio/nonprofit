package handler

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/flintcraftstudio/nonprofit/internal/view"
)

// The public "Carried With Us" site is content-only: each page renders a
// static templ view. They are grouped here rather than split into one file
// per route because none of them parse input or touch the store — they are a
// single feature (the marketing site). Interactive/mutating pages (contact,
// auth) keep their own files.

// About handles GET /about.
func About() http.HandlerFunc {
	return renderPage("about", view.AboutPage())
}

// Podcast handles GET /podcast.
func Podcast() http.HandlerFunc {
	return renderPage("podcast", view.PodcastPage())
}

// Community handles GET /community.
func Community() http.HandlerFunc {
	return renderPage("community", view.CommunityPage())
}

// Resources handles GET /resources.
func Resources() http.HandlerFunc {
	return renderPage("resources", view.ResourcesPage())
}

// Events handles GET /events.
func Events() http.HandlerFunc {
	return renderPage("events", view.EventsPage())
}

// Shop handles GET /shop.
func Shop() http.HandlerFunc {
	return renderPage("shop", view.ShopPage())
}

// Coaching handles GET /coaching.
func Coaching() http.HandlerFunc {
	return renderPage("coaching", view.CoachingPage())
}

// Donate handles GET /donate.
func Donate() http.HandlerFunc {
	return renderPage("donate", view.DonatePage())
}

// renderPage returns a handler that renders a static templ component, logging
// any render error with the page name for context.
func renderPage(name string, page templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := page.Render(r.Context(), w); err != nil {
			slog.Error("render error", "page", name, "err", err)
		}
	}
}
