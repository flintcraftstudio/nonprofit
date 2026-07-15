# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A Firefly Software **Advanced-tier** web app (Go + templ + htmx + Tailwind, SQLite-backed) cloned from the advanced-template starter, being built for **Carried With Us** — a nonprofit supporting bereaved parents through pregnancy and infant loss. The Go module path is `github.com/flintcraftstudio/nonprofit` and `view.SiteName` is `"Carried With Us"`.

Two deep design/architecture docs already exist and are the source of truth — read them before non-trivial work:
- `ARCHITECTURE.md` — layering rules, request lifecycle, and the *why* behind every stack decision.
- `README.md` — full setup, Mage targets, auth model, and middleware config reference.
- `.impeccable.md` — the "Flint & Ember" visual design system (dark/warm, ember accent, Cormorant Garamond + DM Sans). Follow it for any UI work; the `ff-*` Tailwind tokens live in `tailwind/tailwind.config.js`.

## Commands

Everything goes through **Mage** (`magefile.go`). Build logic is Go, not a Makefile.

```bash
mage installtailwind                       # one-time: download the pinned Tailwind standalone CLI
mage seed admin@example.com password       # create DB, run migrations, seed an admin (no registration flow exists)
mage dev                                    # full build (CSS + templ + sqlc + Go) then run the server
mage build                                  # production build → ./bin/server
mage generate                               # templ generate + sqlc generate (regenerate after editing .templ or queries/)
mage migratestatus | migrateup | migratedown
mage createmigration <name>                # scaffold a goose migration
```

Tests are standard Go:
```bash
go test ./...
go test ./internal/middleware/ -run TestCSRF   # single package / single test
```

### Critical: generated code is gitignored
`internal/db/` (sqlc output) and `internal/view/*_templ.go` (templ output) are **not committed**. A fresh clone will not compile until you run `mage generate`. After editing any `queries/*.sql` or `*.templ` file, run `mage generate` (or `mage build`) or the Go code won't see your changes.

## Architecture (the non-obvious parts)

Strict one-directional layering — respect it:
- **`internal/handler/`** — one file per feature. Parse input, call `store`, render `view` components. **No SQL here.**
- **`internal/store/`** — the *only* package that touches the query layer. Wraps sqlc-generated `internal/db` and adds business logic (bcrypt hashing lives here, so plaintext passwords never travel past handler→store).
- **`internal/view/`** — render-only templ components. Never reads env vars; the few globals (`SiteName`, `GtagID`, `PixelID`, `TurnstileSiteKey`) are set once in `main.go` at startup.
- **`internal/middleware/`** — deliberately generic `func(http.Handler) http.Handler` with config structs and **no project imports**, so it can be copied between projects unchanged. Each has a `_test.go`.
- **`internal/session/`** — depends on a `session.Store` interface, not the concrete store.

**Adding a DB query:** write SQL in `queries/*.sql` → `mage generate` → expose it as a method on `Store` in `internal/store/store.go`. Handlers call the `Store` method, never `internal/db` directly.

**Routing & middleware** are wired in `cmd/server/main.go`. The stack wraps inside-out (outermost runs first): Logging → Recovery → RateLimit → CSRF → session.Middleware → mux. `POST /login` gets an *extra* strict per-IP limiter on top of the global one. Protect routes by wrapping in `session.RequireAuth` (see `GET /admin`).

**Migrations are embedded** (`migrations/embed.go`) and run automatically at server/seed startup via goose — a deploy is self-migrating. Mage migrate targets are only for manual dev work.

**htmx pattern:** forms `hx-post` with `hx-swap="outerHTML"`; handlers re-render just the form component with inline errors on failure, and send `HX-Redirect` (alongside a plain `http.Redirect`, so non-JS posts still work). All mutating forms must include `@view.CSRFField()`.

**Auth model:** DB-backed sessions (a `crypto/rand` token is the primary key of a `sessions` row), not JWTs — logout deletes the row for instant revocation. `session.Middleware` never blocks; it only attaches the user to context. Enforcement is separate via `session.RequireAuth`. templ components can call `session.FromContext(ctx)` directly (returns `nil` when unauthenticated — safe on public pages).

## Config & deploy

Config is env-only (`internal/config/`), loaded from `.env` in dev (`cp .env.example .env`). Integrations degrade gracefully when unconfigured (missing Postmark/Turnstile/analytics keys log a warning and disable the feature). **Production is strict:** the server refuses to start without a real ≥32-byte `SESSION_SECRET` (`openssl rand -hex 32`).

Push to `main` → GitHub Actions builds a multi-stage Docker image (runs `mage generate` inside), pushes to GHCR, and redeploys to the Hetzner VPS behind Caddy, pinned to the commit SHA. SQLite lives on the `app-data` volume at `/data/app.db`.

**Tailwind version is pinned in two places** — `magefile.go` and the `Dockerfile`. Keep them in sync when upgrading.
