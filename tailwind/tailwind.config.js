/** @type {import('tailwindcss').Config} */
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: [
    "./internal/view/**/*.templ",
    // Class-emitting Go helpers (e.g. navInitClass in shared.go) live here too;
    // without this glob Tailwind purges rules whose only token is in .go —
    // which silently deleted the base .cw-nav rule and broke the nav.
    "./internal/view/**/*.go",
  ],
  theme: {
    extend: {
      // Short-viewport variant (e.g. landscape phones ~375–560px tall) so the
      // hero can shed its names field and tighten padding before it clips.
      screens: {
        short: { raw: "(max-height: 560px)" },
      },
      colors: {
        ff: {
          // Flint — warm charcoal neutrals (cool stone, faintly warmed)
          dark:    "#121010", // page background
          dark2:   "#1b1714", // glow base / secondary background
          panel:   "#191512", // card surface
          panel2:  "#241e18", // raised surface

          // Ember — the single warm spark accent
          ember:         "#db7b34",
          "ember-hover": "#ec8d44",
          "ember-mid":   "#6b4124", // muted ember: borders, dividers
          "ember-lo":    "#3a2417", // deep ember wash: bg tints, glow cores
          glow:    "#e8b483", // warm-gold light accent: links, focus, inline emphasis

          // Light text — warm whites, never pure white
          paper:   "#f5efe6", // brightest
          cream:   "#e7ddcf",
          moon:    "#d9c4a8", // warm secondary-bright
          ash:     "#948a7c", // muted secondary text
          stone:   "#4d463b", // faint / disabled

          // Hairlines — warm white alpha reads neutral on charcoal
          border:  "rgba(245,239,230,0.07)",
          border2: "rgba(245,239,230,0.13)",
        },

        // "Carried With Us" — the public nonprofit palette. Warm cream light
        // theme lifted by a dawn of muted violet, plum ink, and candle gold.
        // Distinct from the ff-* (Flint & Ember) admin theme on purpose.
        cw: {
          bg:          "#f4efe5", // page cream
          card:        "#f7f2e9", // card / raised surface
          "card-hover":"#faf5ec", // pathway tile hover (lifted cream)
          tile:        "#faf6ef", // donate amount tile resting surface
          sand:        "#efe7d8", // deeper cream (section gradients)
          ink:         "#2c2b3f", // primary headings / plum-black text
          night:       "#24243a", // footer + darkest sections (also theme-color meta)
          violet:      "#55578a", // secondary text + primary button
          "violet-hover":"#494b7a", // primary button hover (deepened violet)
          "violet-lo": "#8a6f8f", // muted violet
          slate:       "#4a4960", // body copy on light
          gold:        "#b98a3e", // links
          "gold-deep": "#7d5c22", // link hover / accent text (text-safe: 5.4:1 on cream)
          glow:        "#e4be83", // candle gold — donate + accents
          "glow-hover":"#d9ad6e", // gold button hover (deepened candle gold)
          clay:        "#c8956c", // warm episode-number / minor accent
          terracotta:  "#a3543f", // eyebrow / label accent (text-safe: 4.7:1 on cream)
          muted:       "#6f645f", // small muted labels (text-safe: 5.0:1 on cream)
          tan:         "#c8b9a3", // warm button borders
          name:        "#f6ead0", // drifting hero remembrance names
          line:        "rgba(85,87,138,0.12)", // violet hairlines
        },
      },
      fontFamily: {
        display: ['"Cormorant Garamond"', ...defaultTheme.fontFamily.serif],
        body:    ['"DM Sans"', ...defaultTheme.fontFamily.sans],
        // Public "Carried With Us" typefaces
        lora:    ['"Lora"', "Georgia", "serif"],
        sans3:   ['"Source Sans 3"', "system-ui", "sans-serif"],
      },
    },
  },
  plugins: [],
}
