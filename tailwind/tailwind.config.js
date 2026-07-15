/** @type {import('tailwindcss').Config} */
const defaultTheme = require("tailwindcss/defaultTheme");

module.exports = {
  content: [
    "./internal/view/**/*.templ",
  ],
  theme: {
    extend: {
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
          sand:        "#efe7d8", // deeper cream (section gradients)
          ink:         "#2c2b3f", // primary headings / plum-black text
          night:       "#24243a", // footer + darkest sections
          violet:      "#55578a", // secondary text + primary button
          "violet-lo": "#8a6f8f", // muted violet
          slate:       "#4a4960", // body copy on light
          gold:        "#b98a3e", // links
          "gold-deep": "#8f6a2c", // link hover / accent text
          glow:        "#e4be83", // candle gold — donate + accents
          terracotta:  "#c27b64", // eyebrow accent
          muted:       "#8a7d78", // small muted labels
          tan:         "#c8b9a3", // warm button borders
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
