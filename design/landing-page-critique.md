# Landing Page Critique — Carried With Us

Design and technical review of the public landing page (`/`, rendered by `internal/view/home.templ`
inside `SiteBase` from `internal/view/site_layout.templ`). Findings are from reading the templ/CSS
source **and** from rendering the built site headlessly at desktop (1440×900), tall, and mobile
(390×844) viewports — several findings below are only visible in the rendered output, not the source.

No code has been changed. Each finding states the problem, why it matters, and how to fix it.

---

## Design Health Score (Nielsen heuristics, 0–4 each)

| # | Heuristic | Score | Key Issue |
|---|-----------|-------|-----------|
| 1 | Visibility of System Status | 1 | The nav is completely invisible on the landing page (see P0-1); active-page state can't be seen |
| 2 | Match System / Real World | 4 | Language is exceptional — compassionate, plain, audience-fluent |
| 3 | User Control and Freedom | 2 | 2.5 s unskippable reveal delay; mobile menu has no Escape-to-close |
| 4 | Consistency and Standards | 3 | Duplicate CTA labels; arbitrary hex values bypassing the `cw-*` tokens |
| 5 | Error Prevention | 3 | Little to prevent on a static page; all links resolve |
| 6 | Recognition Rather Than Recall | 2 | With the nav invisible, everything must be found by scrolling; 9 nav items when working |
| 7 | Flexibility and Efficiency | 2 | Single linear path; no skip link, no way past the intro animation |
| 8 | Aesthetic and Minimalist Design | 3 | Strong, restrained concept; blank strip at top and centered-everything hold it back |
| 9 | Error Recovery | 1 | Unknown URLs soft-404 to the homepage with HTTP 200 (see P1-6); no real 404 page |
| 10 | Help and Documentation | 4 | The crisis-support note in the footer is exemplary contextual help |
| **Total** | | **25/40** | **Acceptable — dragged down hard by two bugs; fixing P0-1 and P1-6 alone lifts this to ~30 (“Good”)** |

## AI-Slop / Anti-Patterns Verdict

**Pass, with caveats.** The core concept — a dawn gradient carrying children's names like
candlelight — is genuinely distinctive, emotionally exact for this audience, and not something a
template produces. The copy is the strongest asset on the page. Nobody would look at the hero and
say "an AI made this."

Tells that remain and should be tightened:

- **Centered-everything**: all four sections are center-aligned, including two long mission
  paragraphs. This is the single most templated-feeling trait of the page.
- **Uniform card grid**: the four pathway tiles are identical in size and structure.
- **Pill buttons everywhere**: every action on the page is the same rounded-full pill; hierarchy
  is carried by color only.

## What's Working

- **The hero concept.** The remembered-names field over a sunrise is the memorable thing the client
  will talk about after the demo. Keep it; fix its execution bugs (below).
- **The copywriting.** "You are not alone in this. You never were." / "plan a goodbye instead of a
  nursery" / "Listen when you're ready" — specific, humane, zero nonprofit boilerplate. The
  footer's "If tonight is heavy" crisis note (988 lifeline) is exactly right for this audience.
- **Reduced-motion handling.** Every animation is silenced under `prefers-reduced-motion`, with
  sensible static fallbacks — rare to see done at all, let alone correctly.
- **Theme discipline.** The `cw-*` public palette is warm and cohesive and stays fully separate
  from the `ff-*` admin theme.

---

## Findings

### P0 — Blocking

**P0-1. The site navigation is invisible and inert on the landing page (Tailwind purged `.cw-nav`).**
- **What**: The built CSS contains `.cw-nav--hero` and `.cw-nav--solid` but **not** the base
  `.cw-nav` rule (`position: fixed`, z-index, transitions). Verified against the served
  `site.css`. Result on `/`: the nav is a static in-flow block, so its `--hero` state paints
  cream text (`#f4efe5`) on the cream body — logo, all 8 links, Donate button, and the mobile
  hamburger are all completely invisible; an empty ~85 px cream bar sits above the hero and pushes
  it down. The scroll-aware solid/transparent transition never happens anywhere on the site, and
  interior pages compensate with `pt-[70px]` for a fixed bar that isn't fixed.
- **Why it matters**: First-time visitors (and the client watching the demo) get a landing page
  with no visible navigation at all. This is the difference between "polished demo" and "broken
  demo."
- **Root cause**: Tailwind v3 only emits `@layer components` rules whose class names appear in
  scanned content. The content glob in `tailwind/tailwind.config.js` is
  `./internal/view/**/*.templ` only, and the bare `cw-nav` token exists solely in
  `internal/view/shared.go` (`navInitClass`). The `--hero`/`--solid` variants survive only by
  luck — they appear inside the Alpine `x-bind:class` string in `site_layout.templ`.
- **Fix**: Add `./internal/view/**/*.go` to the `content` globs (catches this whole class of bug —
  `navInitClass`/`navData` live in Go, and future class-emitting Go helpers will too).
  Alternatively: move the `.cw-nav*` rules out of `@layer components` into plain top-level CSS
  (un-layered rules are never purged), or add `cw-nav` to a `safelist`. The glob fix is the most
  durable.

### P1 — Major (fix before the client sees it)

**P1-1. The "slow sunrise" hero animation is a no-op.**
- **What**: `.cw-hero` sets `background-size: 100% 148%` and animates `background-position` from
  `center 0%` to `center 68%` (`cwSunrise`, 3.6 s). But the hero `<section>` sets the gradient via
  an inline `style="background: linear-gradient(...)"` — the `background` *shorthand* resets
  `background-size` to `auto`, and inline style out-cascades the class. With `auto`, the gradient
  is exactly the element's size, so animating its position by percentage moves it **zero pixels**.
  The marquee animation of the page does nothing (confirmed by comparing screenshots at t=0 and
  t=8 s — identical gradient position).
- **Why it matters**: The one orchestrated page-load moment the design system calls for is dead;
  visitors instead stare at a static dark screen (see P1-3).
- **Fix**: Set the gradient with the `background-image` longhand (inline or, better, move it into
  the `.cw-hero` rule) so the class's `background-size`/`background-position` and the animation
  actually apply.

**P1-2. Hero text lands on the wrong gradient bands — contrast collapses off the 1440×900 happy path.**
- **What**: The gradient's light band is positioned by percentage of section height, while the
  text is positioned by flex layout — the two drift apart across viewports. Rendered results:
  the eyebrow "A MOVEMENT FOR GRIEVING PARENTS & FAMILIES" (`#3a3252`) sits on the dark violet
  zone at ~1.3:1 contrast (near-invisible) at **both** desktop and mobile; on mobile (390×844)
  the first line of the H1 sits in the violet band at roughly **1.5:1** (`#2c2b3f` on ~`#4a4463`);
  and the drifting name "Ruth" collides with the eyebrow text.
- **Why it matters**: The page's headline — its entire message — is unreadable on phones, which is
  where a grieving parent at 2 a.m. is most likely to be.
- **Fix**: Decouple text from gradient luck. Options, roughly in order of preference: (a) give the
  headline block its own guaranteed-light backdrop (a soft radial cream scrim behind it, which
  also reads as "dawn light"); (b) restructure the hero so the light band is anchored to the
  headline's position rather than a fixed 51–68% of section height; (c) at minimum, lighten the
  eyebrow to a cream tone (it currently only works if it lands on the gold band) and give the
  names field a hard bottom margin so names can never touch the copy.

**P1-3. The headline and CTAs are invisible for the first 2.5 seconds.**
- **What**: `.cw-soft-rise` holds the entire headline block (eyebrow, H1, subhead, both CTAs) at
  `opacity: 0` for a 2.5 s delay, then fades over 1.6 s. With the sunrise broken (P1-1), a visitor
  spends 2.5+ s looking at a mostly-black screen with floating names and no explanation. It also
  makes the H1 — the LCP element — paint at ~2.6–4.1 s, which will tank the Lighthouse
  LCP/performance score in any pre-launch audit.
- **Why it matters**: On this site, confusion in the first seconds isn't neutral — the visitor may
  be in crisis. And the demo will be judged partly on PageSpeed numbers.
- **Fix**: Cut the delay to ≤0.8 s and shorten the fade, or start the copy at partial opacity so
  the LCP paints immediately and only the *rise* is animated. Consider playing the full slow
  intro only once per session (sessionStorage flag).

**P1-4. Screen readers announce 14 unexplained children's names before anything else.**
- **What**: The names field precedes the H1 in the DOM with no ARIA treatment. A screen-reader
  user hears "Ellie, Noah, Grace, Amara, Oliver…" with zero context — the sentence explaining that
  these are remembered children comes **after** the CTAs, last. The names are `pointer-events-none`
  but not `aria-hidden`.
- **Why it matters**: For this specific audience, a list of children's names with no framing isn't
  just confusing — it can genuinely distress a bereaved parent using assistive tech. It's also the
  kind of detail a nonprofit client may specifically check.
- **Fix**: `aria-hidden="true"` on the names field, plus a visually-hidden sentence before it
  (e.g. "The names drifting above belong to children who are loved and remembered") — or move the
  existing caption ahead of the names in DOM order and keep it visually positioned below.

**P1-5. Every unknown URL returns the homepage with HTTP 200 (soft 404), and there's no favicon.**
- **What**: `Home()` in `internal/handler/home.go` is registered on `/` and never checks
  `r.URL.Path`, and Go's mux `/` pattern matches everything. Verified: `/definitely-not-a-page`
  → 200 + full homepage; `/favicon.ico` → 200 + **HTML**. There is also no `<link rel="icon">`
  in `SiteBase`.
- **Why it matters**: Search engines see infinite duplicate homepages (soft-404s); typos silently
  "work"; the browser tab shows a blank/broken favicon during the client demo.
- **Fix**: In the handler, `if r.URL.Path != "/" { http.NotFound(w, r); return }` — or register
  the Go 1.22 exact-match pattern `GET /{$}`. Add a designed 404 page (gentle tone — people
  will hit it from dead links in support-group posts) and a real favicon + `<link rel="icon">`.

### P2 — Minor (next pass)

**P2-1. WCAG AA contrast failures on the light sections.**
Measured against `cw-bg` `#f4efe5`:
- Hero caption "The names above belong to children…" — `cw-muted` `#8a7d78` at 13.5 px ≈ **3.5:1**
  (needs 4.5). This is an emotionally load-bearing sentence; it deserves to be readable.
- Eyebrow `cw-terracotta` `#c27b64` at 13 px ≈ **2.9:1**. Fails even the large-text bar.
- Footer copyright `cw-bg/45` on `cw-night` ≈ **3.8:1** at 13 px.
- Pathway CTA `cw-gold-deep` `#8f6a2c` at 14.5 px ≈ **4.3:1** — borderline; nudge darker.
**Fix**: darken the text-use versions of these tokens (decorative uses can keep the current
values); e.g. a terracotta for text around `#a55a42` and a muted around `#6f645f`.

**P2-2. Primary nav carries 9 items and collapses to a hamburger below `lg` (1024 px).**
About, Podcast, Community, Resources, Events, Shop, Coaching, Contact + Donate exceeds the ~5-item
comfort ceiling, and every tablet/small-laptop visitor loses visible navigation entirely.
**Fix**: keep the emotional core visible (About, Podcast, Community, Resources, Events + Donate)
and move Shop/Coaching/Contact into the footer or a "More" group; with 6 items the bar likely fits
at `md` too.

**P2-3. Mobile menu and keyboard access gaps.**
No skip-to-content link anywhere on the site; the hamburger toggles `aria-expanded` (good) but has
no `aria-controls`, Escape doesn't close the panel, and focus isn't moved into/back out of it.
**Fix**: add a skip link in `SiteBase`, `@keydown.escape.window="open = false"`, `aria-controls`,
and focus the first link on open.

**P2-4. No Open Graph / Twitter Card / canonical metadata.**
When the demo URL is texted or Slacked to the client — the most likely first contact — the preview
card is blank. High pitch impact for ~10 lines of head markup.
**Fix**: add `og:title`, `og:description`, `og:image` (a still of the dawn hero would be perfect),
`og:type`, `twitter:card`, and a canonical URL to `SiteBase`.

**P2-5. Google Fonts loaded via CSS `@import`, including the admin fonts, on every public page.**
`tailwind/input.css` imports Cormorant Garamond + DM Sans (admin-only) *and* Lora + Source Sans 3
from `fonts.googleapis.com` in one stylesheet. That's a render-blocking third-party request chain
(CSS → import CSS → font files), it double-loads two families the public site never uses, and it
leaks visitor IPs to Google — which matters more than usual for a grief-support audience (and was
the subject of the German GDPR ruling on hotlinked Google Fonts).
**Fix**: self-host subsetted woff2 files under the static dir with `@font-face` +
`<link rel="preload">`, and drop the admin families from the public bundle.

**P2-6. Every section is center-aligned, including long-form paragraphs.**
The mission section centers two paragraphs of 50–70 words each — centered ragged edges make
multi-line reading measurably harder, and center-everything is the page's most templated trait.
**Fix**: left-align the mission body copy (the serif pull-quote can stay centered), and let one
section break symmetry — e.g. mission text left with the "Read our story" action right.

**P2-7. Duplicate and self-competing CTAs.**
"Listen when you're ready" and "Find others who understand" appear verbatim twice (hero buttons and
pathway tile titles), and the donate section's two buttons ("Support the mission" / "Give monthly")
link to the identical `/donate` URL.
**Fix**: vary the pathway titles (e.g. "The podcast", "The community" as titles with the current
phrases as descriptions — the tiles already have a small label doing this in reverse), and make
"Give monthly" land on the monthly state (`/donate?frequency=monthly` or anchor).

**P2-8. The donation ask carries no trust or impact signal.**
The section is warmly written but gives a donor nothing concrete: no "what your gift does," no
501(c)(3)/tax-deductibility note, no hint of financial stewardship.
**Fix**: one quiet line under the buttons ("Carried With Us is a 501(c)(3) nonprofit — gifts are
tax-deductible" and/or "$25 sends a care package to a newly bereaved family"). For the POC,
placeholder values are fine; the client will notice its presence.

### P3 — Polish

- **P3-1. Arbitrary hex values bypass the `cw-*` tokens** in `home.templ`/`site_layout.templ`:
  `#faf5ec`, `#494b7a`, `#d9ad6e`, `#3a3252`, `#f6ead0`. Add them as tokens (hover variants,
  `cw-name` color) so palette tweaks don't miss stragglers.
- **P3-2. Short-viewport hero squeeze**: `min-h-screen` + a names field with a 160 px floor +
  large fluid type will collide/clip on landscape phones (~500 px tall). Hide or shrink the names
  field under a `max-height` media query.
- **P3-3. Names cluster in the central 1040 px** on wide screens, leaving the hero's outer thirds
  empty (visible at 1440 px); name font sizes are fixed px and don't scale down on mobile. Let the
  field span a viewport percentage and scale sizes with `clamp()`.
- **P3-4. Pathway tiles are four identical cards.** The brand principle is "size things by
  importance" — consider letting Podcast and Community (the two hero CTAs) span wider or taller
  than Resources/Events.
- **P3-5. `pt-[70px]` magic number** on interior `<main>` must stay in sync with actual nav height
  once P0-1 restores fixed positioning — derive both from one value (CSS var) or use
  `scroll-padding-top`.
- **P3-6. htmx is loaded on pages with no `hx-*` attributes** (the landing page uses only Alpine
  for the nav). Fine for a POC; fold into the pre-launch `/audit-js` pass.

---

## Persona Walkthroughs

**A bereaved parent, on a phone, late at night** (primary audience): Meets a dark screen with
floating names and no headline for 2.5 s (P1-3); when the headline appears, its first line is
~1.5:1 contrast on the violet band (P1-2); there is no visible navigation at all (P0-1). The
crisis-support note — the thing this person may most need — is at the very bottom of the page.
⚠️ The single most important user is the one the current bugs hit hardest. (Worth considering: a
quiet "If tonight is heavy →" link higher on the page, not only in the footer.)

**A potential donor evaluating trust** (the demo's business case): Can't find Donate in the nav
because the nav is invisible (P0-1); reaching the donation section, finds warm copy but no
tax-deductibility, impact, or stewardship signal (P2-8); both buttons go to the same place (P2-7).
A shared link to the site shows a blank preview card (P2-4). ⚠️ Every step of the donor path has
friction.

**A friend/professional looking for something to send a grieving parent**: Well served — the
"Take what you need" pathway explicitly addresses them ("something gentle to hand the people who
want to help"). With the nav restored, Resources is one click. ✅ Best-supported persona.

## Cognitive Load Check

7 of 8 checklist items pass — the page has a clear single focus, tight chunking (4 pathway tiles,
2 CTAs per decision point), and good progressive disclosure. The one failure is **minimal
choices**: 9 top-level nav items (P2-2). **Load: low** — this page's problems are execution bugs
and contrast, not overload.

---

## Suggested Fix Order

1. **P0-1** nav CSS purge (one-line config change; unblocks judging everything else) — then re-verify with `/audit`
2. **P1-1 + P1-2 + P1-3** hero: restore the sunrise, guarantee headline contrast, shorten the reveal — `/polish` + `/typeset`
3. **P1-4** names-field ARIA and **P2-3** skip link / menu keyboard access — `/harden` or `/a11y-htmx`
4. **P1-5** soft-404 + 404 page + favicon — `/harden`
5. **P2-1** contrast tokens and **P2-6** alignment — `/typeset` + `/arrange`
6. **P2-4** OG metadata, **P2-5** self-hosted fonts, **P2-8** donation trust line — `/clarify` + `/seo-meta-audit`
7. Remaining P2/P3 as time allows before the client demo — finish with `/polish`
