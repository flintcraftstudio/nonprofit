# Prompt — P3 polish pass on the Carried With Us landing page

> Paste the section below into a fresh Claude Code session at the repo root
> (`github.com/flintcraftstudio/nonprofit`, branch `carried-with-us-site`).
> The P0/P1/P2 findings are already fixed and committed (`075d4d3`); this pass
> is only the **P3 — Polish** items from `design/landing-page-critique.md`.

---

Address the **P3 polish recommendations** in `design/landing-page-critique.md`
for the public landing page. The P0/P1/P2 findings are already done — do **not**
redo or regress them. Read `design/landing-page-critique.md` (the "P3 — Polish"
section and "Suggested Fix Order"), `.impeccable.md` (the "Flint & Ember" +
`cw-*` design system), and `CLAUDE.md` (stack + the critical "generated code is
gitignored" note) before editing.

Stack: Go + templ + htmx + Alpine + Tailwind (v3, standalone CLI via Mage),
SQLite. Relevant files: `internal/view/home.templ`, `internal/view/site_layout.templ`,
`internal/view/shared.go`, `internal/view/donate.templ`, `tailwind/tailwind.config.js`,
`tailwind/input.css`. The public palette is the `cw-*` tokens in the Tailwind
config; the hero/names motion lives in `tailwind/input.css`.

## Do these

**P3-1 — Fold stray hex literals into `cw-*` tokens.**
The critique's specific hex list is **stale** — re-scan the current source
yourself, don't trust the list. Run something like
`grep -rnoE '#[0-9a-fA-F]{3,6}|\[#[0-9a-fA-F]{3,6}\]' internal/view/*.templ`
and, for each arbitrary literal used as a color (hover variants like
`hover:bg-[#494b7a]` / `hover:bg-[#d9ad6e]` / `hover:bg-[#faf5ec]`, the
`cw-name` color `#f6ead0`, the amount-tile bg `#faf6ef`, etc.), add a named
`cw-*` token in `tailwind/tailwind.config.js` and reference it, so a future
palette tweak can't miss a straggler. Skip literals inside `rgba()` decorative
glows and the SVG assets. Note `#24243a` in the `theme-color` meta already
equals `cw-night` — use the token there too if trivial. Keep the rendered
colors identical (this is a refactor, not a restyle).

**P3-2 — Fix the short-viewport hero squeeze.**
On landscape phones (~500px tall) `min-h-screen` + the names field's
`clamp(160px,26vh,280px)` floor + large fluid type collide/clip. Under a short
viewport (e.g. `@media (max-height: 560px)` or a Tailwind arbitrary variant),
shrink or hide the names field and tighten the hero's vertical padding so the
headline + CTAs never clip. Verify at 812×375 (landscape phone).

**P3-3 — Let the names field breathe on wide screens and scale down on mobile.**
The field is capped at `max-w-[1040px]` so on 1440px+ the outer thirds sit
empty, and the per-name sizes are fixed px (`Size` in `HeroNames`) that don't
scale on small screens. Let the field span a viewport percentage (wider cap or
`vw`-based width) and drive the name sizes with `clamp()` so they shrink on
mobile and spread on desktop. `heroNameStyle` in `shared.go` builds the inline
`font-size` — switch it to a `clamp()` expression (keep the per-name `Size` as
the preferred/mid value). Preserve the reduced-motion static field.

**P3-4 — Break the uniform pathway grid ("size things by importance").**
The four pathway tiles are identical. Per the brand principle, let the two that
are also the hero CTAs — **The podcast** and **The community** — carry more
weight than Resources/Events (e.g. span wider or taller on `sm+`, or a larger
title). Keep it tasteful and keep the grid responsive (single column on mobile).
This is a design-judgment item; aim for gentle emphasis, not a busy layout.

**P3-5 — Kill the `pt-[70px]` magic number.**
Interior `<main>` uses `pt-[70px]` (`site_layout.templ`) to clear the fixed nav,
which must stay in sync with the actual solid-nav height (padding `0.85rem`
top+bottom). Derive both from one source: define a `--cw-nav-h` CSS var (or set
the nav height explicitly) and use it for the main offset, and/or add
`scroll-padding-top` so in-page anchors clear the nav too. Verify the interior
pages (`/about`, `/podcast`, …) have no gap or overlap under the nav.

**P3-6 — htmx on pages with no `hx-*` (optional / defer).**
The landing page loads `htmx.min.js` but only uses Alpine. The critique says
this is fine for the POC and belongs in the pre-launch `/audit-js` pass — so
**only** note it (leave htmx in place unless you can confirm no public page uses
`hx-*`). Do not remove it speculatively.

## Constraints & verification

- Use the `cw-*` public theme only; do not touch the `ff-*` admin theme.
- After editing any `.templ` or the Tailwind config/CSS, regenerate:
  `mage generate` (templ+sqlc) and `mage buildCSS`, then `go build ./...`.
  Generated `*_templ.go` / `internal/db` are gitignored — they won't compile
  from a clean tree until generated.
- **Reduced-motion parity:** anything you change in the hero/names must keep the
  `@media (prefers-reduced-motion: reduce)` block in `tailwind/input.css`
  correct (names rest static at 0.62 opacity; hero rests fully bloomed; no
  clipping). Don't reintroduce autoplaying motion without a reduced-motion rest
  state.
- **Verify by rendering**, not just source: run the server (`mage dev` or a
  built binary), and screenshot the home + an interior page at desktop
  (1440×900), mobile portrait (390×844), landscape phone (812×375), and with
  reduced motion emulated. Confirm no regressions to the P0–P2 fixes (visible
  nav, headline contrast, real 404, self-hosted fonts).
- Keep changes scoped to P3. Don't restyle beyond the items above.

When done, summarize what changed per item and flag P3-6 as deferred to
`/audit-js`.
