package view

import (
	"fmt"
	"math"
	"time"
)

// SiteName is the display name used in templates. Override per-project.
const SiteName = "Carried With Us"

// CwNight is the darkest public surface, used for the browser theme-color meta
// (a raw HTML attribute, so it can't reference the Tailwind class). Keep in
// sync with the cw.night token in tailwind/tailwind.config.js.
const CwNight = "#24243a"

// Tracking IDs and Turnstile site key, set once at startup from config.
var (
	PixelID          string
	GtagID           string
	TurnstileSiteKey string
	// SiteURL is the canonical origin (no trailing slash), used to build
	// absolute canonical + Open Graph URLs. Empty in dev; when empty the
	// templates fall back to root-relative URLs.
	SiteURL string
)

// CanonicalURL returns the absolute URL for a path when SiteURL is configured,
// otherwise the root-relative path itself (a valid relative canonical).
func CanonicalURL(path string) string {
	if SiteURL == "" {
		return path
	}
	return SiteURL + path
}

// AssetURL returns an absolute URL for a static asset when SiteURL is set
// (Open Graph images should be absolute), else the root-relative path.
func AssetURL(path string) string {
	if SiteURL == "" {
		return path
	}
	return SiteURL + path
}

// Year returns the current year for copyright notices.
func Year() int {
	return time.Now().Year()
}

// NavItem is a single primary navigation link on the public site.
type NavItem struct {
	Href  string
	Label string
}

// NavItems are the primary navigation links for "Carried With Us", in order —
// the emotional core, kept to five so the desktop bar stays scannable. Donate
// is rendered separately as an accented button.
var NavItems = []NavItem{
	{"/about", "About"},
	{"/podcast", "Podcast"},
	{"/community", "Community"},
	{"/resources", "Resources"},
	{"/events", "Events"},
}

// SecondaryNavItems are the supporting links kept out of the primary desktop
// bar (which would otherwise be overcrowded). They still appear in the footer
// sitemap and the mobile menu, so nothing becomes unreachable.
var SecondaryNavItems = []NavItem{
	{"/shop", "Shop"},
	{"/coaching", "Coaching"},
	{"/contact", "Contact"},
}

// navInitClass is the server-rendered starting state for the fixed nav:
// transparent over the home hero, solid cream everywhere else. Alpine takes
// over on scroll, so this only governs the first paint (avoids a flash).
func navInitClass(isHome bool) string {
	if isHome {
		return "cw-nav cw-nav--hero"
	}
	return "cw-nav cw-nav--solid"
}

// navData is the Alpine state object for the nav. `home` freezes whether this
// page is allowed to show the transparent hero variant at all.
func navData(isHome bool) string {
	if isHome {
		return "{ top: true, open: false, home: true }"
	}
	return "{ top: true, open: false, home: false }"
}

// HeroName is one child's name drifting in the home hero, remembered like
// candlelight. Positions are percentages within the names field.
type HeroName struct {
	Name  string
	Left  string
	Top   string
	Size  int
	Delay string
	Dur   string
}

// HeroNames are placeholder remembrance names for the demo hero. In a real
// deployment these would be drawn from families who choose to add a name.
var HeroNames = []HeroName{
	{"Ellie", "11%", "6%", 22, "0s", "7.5s"},
	{"Noah", "27%", "34%", 19, "1.3s", "8.5s"},
	{"Grace", "41%", "12%", 26, "2.2s", "7s"},
	{"Amara", "57%", "40%", 20, "0.6s", "9s"},
	{"Oliver", "72%", "9%", 23, "3s", "8s"},
	{"Mia", "86%", "36%", 18, "1.8s", "7.5s"},
	{"Leo", "17%", "62%", 20, "2.6s", "8.5s"},
	{"Ruth", "33%", "74%", 24, "0.9s", "7s"},
	{"Theo", "49%", "64%", 19, "3.4s", "9s"},
	{"June", "64%", "76%", 22, "1.5s", "8s"},
	{"Isaac", "79%", "63%", 20, "2.3s", "7.5s"},
	{"Hazel", "7%", "38%", 18, "3.1s", "8.5s"},
	{"Nora", "90%", "60%", 19, "0.4s", "7s"},
	{"Clara", "46%", "90%", 21, "2s", "9s"},
}

// heroNameStyle builds the absolute-position + animation inline style for one
// drifting hero name. The font-size is a clamp() around the per-name Size (the
// preferred/mid value): it shrinks toward ~0.62× on narrow phones and grows
// toward ~1.22× on wide screens, tracking the viewport-percentage width of the
// names field so the remembrance field scales instead of clustering.
func heroNameStyle(n HeroName) string {
	minPx := int(math.Round(float64(n.Size) * 0.62))
	maxPx := int(math.Round(float64(n.Size) * 1.22))
	prefVw := float64(n.Size) / 10.5 // ≈ Size px around a ~1050px-wide viewport
	return fmt.Sprintf(
		"left:%s;top:%s;font-size:clamp(%dpx,%.2fvw,%dpx);animation-delay:%s;animation-duration:%s",
		n.Left, n.Top, minPx, prefVw, maxPx, n.Delay, n.Dur,
	)
}

// The slices below are placeholder demo content for the proof-of-concept. In a
// real deployment these would come from the store; they are kept here so the
// POC has no database dependency for public content.

// Warm placeholder image gradients (stand in for real photography/art).
const (
	gradRose   = "linear-gradient(150deg,#f0d9c4,#e7bfa0 45%,#d99f8f 100%)"
	gradViolet = "linear-gradient(150deg,#efdcc2,#b98f88 55%,#7d6c82 100%)"
	gradGold   = "linear-gradient(150deg,#f4e6cf,#e6c08a 60%,#d9a86f 100%)"
	gradDusk   = "linear-gradient(150deg,#e9c9b6,#c98f92 50%,#7d6c82 100%)"
)

// Episode is one podcast episode in the demo list.
type Episode struct {
	Num, Title, Desc, Len string
}

var Episodes = []Episode{
	{"23", "When Friends Don’t Know What to Say", "Why the silence hurts, and how to gently ask for what you need.", "39 min"},
	{"22", "Marking Birthdays", "Small rituals for the days that will always be theirs.", "42 min"},
	{"21", "Grieving as a Couple", "Two people, one loss, and two very different ways of carrying it.", "51 min"},
	{"20", "Telling Their Siblings", "Honest, age-appropriate words for the children who are still here.", "44 min"},
	{"19", "“How Many Children Do You Have?”", "Answering a simple question that never feels simple.", "36 min"},
	{"18", "Holidays Without Them", "Making room for both the empty chair and the joy that remains.", "49 min"},
	{"17", "Finding the Others", "How community changes grief, and why you don’t have to do this alone.", "40 min"},
}

// Product is one item in the demo shop.
type Product struct {
	Name, Desc, Price, Tag, Grad string
}

var Products = []Product{
	{"Remembrance Candle", "A candle for the nights you want to remember them together.", "$32", "Remembrance", gradGold},
	{"Birthstone Necklace", "Their birthstone, close to your heart, wherever you go.", "$88", "Jewelry", gradDusk},
	{"“Carried With Us” Sweatshirt", "Soft enough for the hard days — a quiet way to feel held.", "$58", "Apparel", gradRose},
	{"The Grief & Hope Journal", "Space for the words that are hard to say out loud.", "$26", "Journals", gradViolet},
	{"Forever Loved Tee", "Their name matters. So does saying it, out loud and often.", "$34", "Apparel", gradRose},
	{"Remembrance Ornament", "A place for them at every gathering, every single year.", "$28", "Remembrance", gradGold},
	{"First Light Candle Set", "Three candles — for the long nights, and the mornings after.", "$46", "Remembrance", gradDusk},
	{"Name Bracelet", "Their name, in gold, always with you.", "$72", "Jewelry", gradViolet},
}

// Resource is one free download on the resources page.
type Resource struct {
	Title, Desc string
}

var Resources = []Resource{
	{"A Guide to the First Weeks", "What to expect, what to say no to, and how to be gentle with yourself."},
	{"Journaling Prompts", "Thirty quiet prompts for the words that are hard to say out loud."},
	{"Supporting a Grieving Friend", "For the people who love you and don't know what to do. Hand them this."},
	{"The First-Year Companion", "A month-by-month companion through birthdays, holidays, and hard dates."},
}

// Event is one gathering on the events page.
type Event struct {
	Meta, Title, Desc, Grad string
}

var Events = []Event{
	{"Retreat · Oct 17–19, 2026 · Asheville, NC", "First Light Retreat", "A slow weekend in the mountains for bereaved parents — quiet mornings, shared meals, and evenings by the fire. Space is limited and held with care.", "linear-gradient(150deg,#f0d9c4,#e7bfa0 50%,#8f7f9e 100%)"},
	{"Conference · March 2027 · Online & in person", "Carried With Us Conference", "A day of honest talks, small workshops, and a shared time of remembrance. Join us in the room or from your own kitchen table.", "linear-gradient(150deg,#f4e6cf,#e6c08a 55%,#d99f8f 100%)"},
}

// PodcastPlatform is one place to subscribe.
type PodcastPlatform struct {
	Name, Href string
}

var PodcastPlatforms = []PodcastPlatform{
	{"Apple Podcasts", "#"},
	{"Spotify", "#"},
	{"YouTube", "#"},
	{"RSS", "#"},
}
