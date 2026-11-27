---
name: I fucking hate snaps.
description: A whack-a-mole board where you finally get to hit a snap, and one real complaint per page.
colors:
  ground: "#fdf7ec"
  night: "#222a3a"
  night-ink: "#d9d3c8"
  ink: "#2e2925"
  ink-soft: "#4a433d"
  ink-muted: "#6b625b"
  plate-ink: "#0b0b0b"
  red: "#d8433a"
  hole-deep: "#050403"
  hole-dark: "#120e0b"
  hole-wall: "#33281f"
  hole-lit: "#5e4b3c"
  dirt: "#e9dac1"
  hole-shadow: "rgba(90, 60, 30, 0.16)"
  ground-dark: "#24211e"
  ink-dark: "#e2dacd"
  ink-soft-dark: "#c4bcaf"
  ink-muted-dark: "#9d9488"
  red-dark: "#e07a6f"
  night-dark: "#1b1c20"
  night-ink-dark: "#c6c0b6"
  hole-deep-dark: "#000000"
  hole-dark-dark: "#070605"
  hole-wall-dark: "#15110e"
  hole-lip-dark: "rgba(255, 238, 214, 0.07)"
typography:
  display:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "max(calc(140 * var(--px)), 3rem)"
    fontWeight: 800
    lineHeight: 1
    letterSpacing: "-0.01em"
  headline:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "max(calc(99 * var(--px)), 2.25rem)"
    fontWeight: 800
    lineHeight: 1
  pull:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "clamp(3.125rem, 7.2vw + 1rem, 6.5rem)"
    fontWeight: 800
    lineHeight: 1.02
    letterSpacing: "-0.02em"
  reading:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "clamp(1.125rem, 1.8vw, 1.625rem)"
    fontWeight: 400
    lineHeight: 1.5
  masthead:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "clamp(1.125rem, 1.7vw, 1.5rem)"
    fontWeight: 800
  button:
    fontFamily: "Figtree, sans-serif"
    fontSize: "1.1875rem"
    fontWeight: 800
    lineHeight: 1
  body:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "clamp(0.875rem, 1.25vw, 1.0625rem)"
    fontWeight: 400
    lineHeight: 1.5
  meta:
    fontFamily: "Figtree, Arial Black, sans-serif"
    fontSize: "1rem"
    fontWeight: 400
rounded:
  focus: "2px"
  oval: "50%"
  pill: "999px"
spacing:
  gutter: "1rem"
  focus-offset: "6px"
components:
  hole-pit:
    rounded: "{rounded.oval}"
  report-band:
    backgroundColor: "{colors.night}"
    textColor: "{colors.night-ink}"
    typography: "{typography.body}"
    padding: "1.15em clamp(1rem, 4vw, 3.5rem) 1.25em"
  link-report:
    textColor: "{colors.night-ink}"
    typography: "{typography.body}"
  pill:
    backgroundColor: "{colors.ink}"
    textColor: "{colors.ground}"
    typography: "{typography.button}"
    rounded: "{rounded.pill}"
    padding: "0.85em 1.5em"
  pill-hover:
    backgroundColor: "{colors.ink-soft}"
  link-action:
    textColor: "{colors.ink}"
    typography: "{typography.button}"
---

# Design System: I fucking hate snaps.

## Overview

**Creative North Star: "The Whack-a-Mole Board"**

The landing page is a fairground game drawn as flat clean vector: a warm off-white ground, heavy warm-charcoal grotesque type, three dark oval holes with real depth, and a smug cardboard box labelled "SNAP" that pops out of them. It is not a hero with a button; the box is the primary action. Everything is soft charcoal on off-white except a muted night-blue band along the bottom edge of every page, a red underline on hovered links is the only colour, and nothing but the holes has a gradient or shadow; nothing has texture or gloss; the one worn edge in the world is the ink stamp printed on the box lid.

Density is low and the column is centred: headline, board, call to action, and one closing sentence on a full-width night band on the bottom edge, nothing else. The page never scrolls. The whole composition is drawn on a 1536 x 1024 comp and scaled as one unit, so the proportions of the approved comp hold at every desktop size. Motion belongs to the game alone: the box bounces out of its hole with a stretch and squash, wobbles smugly while it waits and only moves when the mallet reaches it, ducking back down with a little hop and dodging to another hole, follows the mallet with its eyes, flinches and gets more battered with each of its three hits, squashes flat on the last, and under reduced motion appears and disappears instantly. With a mouse, the pointer is a drawn flat wooden mallet that swings down on press, and the hit lands with a jolt of the board and a synthesised bonk.

The quote page (`/<n>/`) and the 404 live in the same world as reading pages, "Punchline first": a small masthead linking home, the highlighted punchline pulled out huge in quotes the way people would quote it, the whole complaint below at reading size, the permanent link in muted ink, one ink pill to copy or share it, plain links to add your own rant and to another random opinion (the previous quote is linked only as `rel="prev"` in the head), and the same night footer. No game, no plates, no motion beyond the browser's own. All pages share one stylesheet (`site.css`) for fonts, palette, the page frame and the footer band; the landing page adds the game on top. Every page carries a link preview (`og.jpg`, 1200 x 630, cut from the approved landing comp), a canonical URL and the box favicon.

The world has two first-class renditions, light and dark. Dark is the same world with the lights down: soft warm charcoal paper, not near-black, warm light ink, a quiet neutral footer band, a lighter red so a hover underline still shows, and holes shaded their own way so they read as sunk, not as mounds: darker than the ground everywhere, with no dirt ring and no ground shadow. The drawn plates are dimmed slightly to sit in the dim scene; the pupils do not change. Pages follow the system preference, and a small round picker in the top-right corner lets the visitor choose system, light or dark.

**Key Characteristics:**
- Flat clean vector on a warm off-white ground (soft warm charcoal in dark); soft warm ink above a muted night-blue footer band (a quiet neutral in dark); contrast deliberately softened.
- Light and dark are both first-class: system preference by default, the visitor's pick when made.
- One red, used only as a link's hover underline.
- Figtree at two weights: 800 for display lines, the punchline, the masthead and the button; 400 for reading and small text.
- Oval holes with depth (in light a lit far wall falling into darkness, a dirt ring and a soft ground shadow; in dark sunk below the ground, darkest under the far rim); everything else is flat.
- Two raster plates in the same flat style: the smug cardboard "SNAP" box and the wooden mallet.
- Reading pages: punchline first, the whole quote at reading size, one ink pill for the page's main action.
- Motion is the game (landing only): bouncy pop, taunting wobble, duck, dodge, eyes on the mallet, ouch, squash, mallet strike, board thud; under reduced motion the snap moves instantly and nothing bounces, wobbles, slides, squashes or jolts.

## Colors

Two paper-and-ink pairs (the page and the night band), two softer inks for reading, and one accent that appears only under links. Every role has a light value (the plain token) and a dark value (the `-dark` token); the role descriptions hold in both.

### Primary
- **Underline Red** (`red`): only ever the underline of a hovered link: the action links on reading pages and the footer sentence links (1.5px, offset 0.22em, continuous with no skip-ink), and the masthead link. Nothing is red at rest. It never fills an area, colours text, or marks anything else.

### Neutral
- **Warm Paper** (`ground`): the page background and selection text colour. Set on `html` so overscroll stays the same colour.
- **Night** (`night`): a muted dark slate blue; the full-width band on the bottom edge that carries the closing sentence. It is the only filled area on the page besides the holes.
- **Night Ink** (`night-ink`): a warm off-white for text on the night band, its resting link hairline and its focus outlines.
- **Warm Charcoal** (`ink`): display text, links and the punchline highlight above the band, the ink pill's fill, focus outlines on the page, and the selection background. Softer than black on purpose.
- **Soft Ink** (`ink-soft`): reading text (the full quote, the 404 lede) and the ink pill on hover.
- **Muted Ink** (`ink-muted`): meta only, the "Opinion No. N" line with its permanent link.
- **Plate Ink** (`plate-ink`): pure black, only for the code-drawn pupils, so they match the plates' black outlines.

### Hole Material
Used only inside the holes' depth treatment (see Elevation & Depth), never as UI colour.
- **Hole Deep** (`hole-deep`), **Hole Dark** (`hole-dark`), **Hole Wall** (`hole-wall`), **Hole Lit** (`hole-lit`): the stops of the light hole fill, from the dark near rim (0%, 40%) up to the lit far wall (72%, 100%).
- **Dirt** (`dirt`): the ring of turned-up dirt round each hole, in light only.
- **Hole Shadow** (`hole-shadow`): the soft shadow each hole throws on the ground, in light only.
- **Dark hole** (`hole-deep-dark`, `hole-dark-dark`, `hole-wall-dark`): the stops of the dark hole fill, darkest under the far rim; **Hole Lip** (`hole-lip-dark`): a faint warm light on the near lip in dark.

### Dark
The lights down, not night: soft warm charcoal, never near-black, with contrast kept gentle. Applied when the system prefers dark and the visitor has not picked light, or when the visitor picked dark. `color-scheme` follows.
- **Warm Charcoal Paper** (`ground-dark`), **Warm Light Ink** (`ink-dark`), **Soft Light Ink** (`ink-soft-dark`), **Muted Light Ink** (`ink-muted-dark`): the page and its inks.
- **Lifted Red** (`red-dark`): the hover underline, lighter so it shows on dark.
- **Quiet Night** (`night-dark`), a neutral near the ground, no longer blue, and **Night Ink** (`night-ink-dark`): the footer band.
- **Holes:** their own dark shading (see Elevation & Depth and Hole Material). `dirt` and `hole-shadow` are light-only tokens; the dark holes have no dirt ring or ground shadow.
- **Plate dim:** every `img` (the box and mallet plates) takes `filter: brightness(0.88) saturate(0.9)` in dark (`--plate-dim`; none in light), so the drawn plates sit in the dim scene.
- **Unchanged:** plate ink. The ink pill flips with the inks into a light pill with dark text.

### Named Rules
**The One Red Rule.** Red is only ever a link's hover underline; at rest the page has no red. Red on text, fills, icons, a resting underline or anything else breaks the world.

**The Black Ink Rule.** Every glyph is one of the inks on paper, except on the night band (night ink on night) and the ink pill (paper on ink). Emphasis comes from size, weight and an ink underline, never from colour.

The cardboard tan, tape and outline colours of the box live inside the raster plate. They are a material, not palette tokens, and do not license tan or orange anywhere in CSS.

## Typography

**Display Font:** Figtree 800 (with Arial Black, sans-serif)
**Body Font:** Figtree 400 (same stack)

**Character:** One geometric grotesque at a heavy weight and its regular weight. The heavy weight carries the joke at billboard size; the regular weight carries reading text and the quiet lines. Both are self-hosted Latin-subset woff2 files under the OFL; the 800 is preloaded.

### Hierarchy
- **Display** (800, 140 comp px with a 3rem floor; clamp(3.25rem, 11vw, 7rem) in the compact layout; line-height 1; -0.01em; balanced wrap): the headline "I fucking hate snaps." only.
- **Headline** (800, 99 comp px with a 2.25rem floor; clamp(2.5rem, 8vw, 4.5rem) compact; line-height 1): the call to action under the board, "Whack a snap."
- **Pull** (800, clamp(3.125rem, 7.2vw + 1rem, 6.5rem), line-height 1.02, -0.02em, 13em max, balanced): the punchline on a quote page, in curly quotes, and "No opinion here." on the 404.
- **Reading** (400, clamp(1.125rem, 1.8vw, 1.625rem), line-height 1.5, 34em measure, pretty wrap): the full quote in soft ink. A quote with no highlight is shown "solo" instead: ink, clamp(1.75rem, 3.8vw, 3rem), line-height 1.3, 22em measure. The 404 lede is soft ink at clamp(1.25rem, 2.4vw, 2rem).
- **Masthead** (800, clamp(1.125rem, 1.7vw, 1.5rem)): "I fucking hate snaps." at the top of reading pages, linking home.
- **Button** (800, 1.1875rem, line-height 1): the ink pill label; the action links sit beside it at the same size in 400.
- **Body** (400, clamp(0.875rem, 1.25vw, 1.0625rem); on the landing page 23 comp px with a 0.9375rem floor and clamp(0.875rem, 2.4vw, 1.0625rem) compact; line-height 1.5; left-aligned, pretty wrap): the closing sentence on the band.
- **Meta** (400, 1rem, muted ink): "Opinion No. N · ifuckinghatesnaps.com/N/"; the link is selectable as one unit.

The headline and call to action trim their text box to cap height and alphabetic baseline (`text-box: trim-both cap alphabetic`), so the measured gaps between them are optical, glyph to glyph.

### Named Rules
**The Two Weights Rule.** Figtree 800 (display, punchline, masthead, button) and 400 (reading and small text) are the only faces. No italics, no other weights, no second family.

**The Curly Quotes Rule.** Quote text is set with typographic punctuation: straight quotes are converted to “ ” ‘ ’ when the quotes load.

## Layout

Every page is a full-height flex column: the page's own content fills what the night footer band leaves and is centred in it, in a single centred column with a 1rem side gutter, so short pages never scroll.

**Reading pages** (quote, 404): the masthead sits 1.75rem from the top; the content column has clamp(2rem, 6svh, 4.5rem) padding top and bottom. The punchline, the quote, the meta line and the actions stack with clamp(1.5rem, 4.5svh, 2.75rem) between the big steps and 1.1rem above the meta. Actions wrap centred with 1rem by 1.75rem gaps; below 40rem they become a two-column grid (1.1rem by 1.75rem gaps) with the share pill on its own line and the two action links paired below it. No comp unit here: reading pages size with viewport clamps.

**The Comp Unit Rule.** On the landing page, desktop measurements are written in comp pixels of the 1536 x 1024 design: one comp pixel is `min(1vw, 1.5svh) / 15.36`, so the layout scales with whichever of width or height is tighter. At 1536 x 1024 the rhythm is: 48 top padding, headline, 122 to the board, board 1242 x 248, 104 to the call to action. The page never scrolls: the body is a full-height flex column, `main` takes the space the footer band leaves (48 comp px padding top and bottom) and centres the game in it, and the band sits on the bottom edge on every size.

**Compact layout** (width up to 40rem, or portrait up to 64rem): the holes form a triangle, left and right on top and the middle one centred below, with an 80-unit row gap on an 820-unit-wide board. One board unit is `min((100vw - 2rem) / 820, (100svh - 25rem) / 640)`, so the board is capped by the height too and the page never scrolls (the snap is about 155px wide on a 390px phone). The type switches to viewport clamps, and `main` has 2rem padding top and bottom.

## Elevation & Depth

The world is flat except for the holes. Nothing else has a box-shadow, drop-shadow, gradient, blur, texture or gloss. Each hole is shaded as if looking down into it: a radial fill whose far wall is lit and falls into darkness toward the near rim, an inset shadow along the near rim where it overhangs, a 4 comp px ring of turned-up dirt round the edge, and a soft warm ground shadow under it. In dark a lit far wall on a dark ground reads as a mound, so the hole is shaded the other way: a radial fill darker than the ground everywhere and darkest under the far rim (from the top centre outward), an inset shadow along the far rim's overhang (12 by 14 comp px, black at 75%), a faint warm light on the near lip (2 comp px), and no dirt ring or ground shadow. Everything else gets its depth from overlap. The box sits in a clipped window whose bottom edge follows the front rim of its oval, so it rises out of and sinks into the hole; the flattened box, once hit, sits over the rim outside that clip.

### Named Rules
**The Flat World Rule.** Only the holes are shaded. Anything else that needs to look lifted or sunk does it with overlap and clipping against flat shapes, never with shading.

## Shapes

Flat ellipses and the box's own drawn silhouette. The holes are shaded ovals (border-radius 50%) with a dirt ring: two outer holes 350 x 108 comp px and a slightly wider, flatter middle hole 394 x 103. No cards, borders, dividers or containers; the night band is square-cornered, has no rule, and runs edge to edge. The rounded controls are the ink pill (fully rounded, 999px) and the round 44px theme picker. The only small radius is the 2px corner on link focus outlines.

## Components

### Board and Holes (signature)
Three shaded oval holes in a row (a triangle in the compact layout; see Layout and Elevation & Depth), spread across the board's width and aligned to its bottom edge. The middle hole holds the box at rest. Each hole is a pit (the oval), a window that clips everything above the pit's front rim, and a hidden flattened-box image for the hit state.

### The Snap (signature)
The smug cardboard "SNAP" box plate, a link to `/random`. Its label counts down the hits: "Whack a snap: N hits left to open a random opinion" (without JS: "Whack a snap: open a random opinion", and one click opens a quote).
- **Plate:** a smug box with half-lidded eyes (empty whites; the pupils are drawn in code) and a thick one-sided smirk. Beside the tape on the lid, on both the upright and the flattened plate, sits a worn black-ink stamp: a parody folded-paper bird of four chunky triangles, with no circle, text or colour. It is a raster illustration served as webp with a png fallback (800 x 640) and a flattened variant with x eyes (1000 x 400). A 64px favicon is cut from the upright plate.
- **Battered stages:** two more upright plates replace every snap's plate after hits 1 and 2. Stage one has a dent, half-peeled tape and an annoyed wavy mouth; stage two has a crushed corner, a torn flap, crossed pale sticking plasters on the side panel and a gritted grimace. Both keep the eye whites, lids, SNAP and the lid stamp exactly in place, so the drawn pupils still align. Source plates live in `assets/plates/` with their generation provenance embedded. Every shipped raster carries its provenance: PNGs as PNG tEXt, `og.jpg` as a JPEG comment segment, and only WebP files as a `.json` note next to them.
- **Pop:** a 560ms bounce anchored at the bottom centre: it stretches out of the hole (0.88 x 1.18), overshoots to 17% above rest (0.92 x 1.1), lands squashed (1.14 x 0.84), rebounds to 6% above rest, settles through 1.04 x 0.96 to rest. Each segment carries its own easing. The clip window above the hole reaches 140 comp px up so the overshoot fits.
- **Taunt:** after the pop it rocks ±3° about its bottom centre, 1.3s ease-in-out, for as long as it is up; once tired, the taunt slows to 2.6s.
- **Duck:** a 320ms drop: a small hop to 5% above rest (0.96 x 1.05), then an accelerating slide 110% down into the hole (1.02 x 0.98). Snaps the server renders down skip it on load.
- **Dodge:** the snap never moves on its own; it stays up and taunts. When a mouse pointer enters the up snap, it is back in the hole 0.45 to 0.75 seconds later (it starts its 320ms duck that much earlier; under reduced motion the duck is instant, so the full delay applies), and after a 450 to 650ms gap pops from a different hole. If it comes up under a resting mallet, it checks the last pointer position 220ms into its pop (at once under reduced motion) and dodges again without the mouse moving. Focus, a hit, or hiding the page cancels a pending dodge, and a snap caught between holes comes back up.
- **Touch:** on devices without hover the snap hops by itself every 1.6 to 2.6 seconds, plus fatigue x 300ms, until it is tired.
- **Tiring:** every dodge, every touch self-hop and every mouse miss on the board adds 1 fatigue, and chase time adds 1 per 1.6 seconds since the first dodge. Each dodge is delayed by fatigue x 160ms. At an effort of 6 the snap is tired: it stops dodging and fleeing and just taunts. A hit removes 2 fatigue and restarts the chase clock. In practice it gives up after about 4 dodges or 9 seconds.
- **Hit:** with the mallet, the hit is decided where the head lands, 70ms after the press, at the current pointer position. Any snap whose drawn box shows above its hole's rim there is hit, whether it is popping up, flinching or sliding down: the drawn box is the plate inset 7.5% left and right and 6% from the top, down to 92% of the plate or the rim, whichever is higher on screen. The plate's transparent margin and the link's own hit area do not count. Each press counts once, and a click within 1s of a mouse press is ignored. Keyboard and touch whack on click. Modified clicks open normally. Every hit thuds the board (200ms ease-out, down 6px, up 2px, rest).
- **Three lives:** hits 1 and 2 swap in the next battered stage and play the hit sound. A snap hit while up plays an `ouch` flinch (260ms: squashed 1.18 x 0.78, springs to 0.94 x 1.07 at 4% above rest, settles), then the taunt carries on; a snap hit while sliding down takes the damage and keeps going. Hit 3 replaces whichever snap was hit with the flattened plate (its drawn eyes hidden), which squashes in over 260ms (tall and narrow, overshoots wide and short, settles), ducks any other snap that is up, plays the kill sound, and navigates after 650ms.
- **Flee:** after a mouse hit on a snap that is up and not tired, it runs at once, hopping 260ms later. A duck always wins over a flinch still playing, so a fleeing snap always goes down. Keyboard and touch hits do not make it flee.
- **Announcements:** a visually hidden status line reads out "Ouch. 2 hits left.", "Ouch. 1 hit left." and "Got it. Opening a random opinion."
- **Back:** returning with the Back button reloads the page for a fresh game.
- **Keyboard:** exactly one snap is in the tab order, the one that is up or about to come up; the others are unfocusable and hidden from assistive tech. Focusing it brings it up at once and pauses the game until blur.
- **Focus:** the box shows no outline of its own; its hole's oval gets a 3px ink outline at 6px offset.
- **Eyes:** the pupils are flat plate-ink dots, each clipped to its eye white under the lid. At rest they keep the smug sideways glance (x 0.6, y 0.3 on a -1 to 1 range); with a mouse, the up snap's pupils follow the pointer, eased over 90ms.
- **Reduced motion:** the game keeps running, but with no pop, taunt, duck, ouch, transitions, pupil easing, squash or thud, so the snap appears and disappears instantly; a mallet hit lands at once on the press, a fleeing snap hops after 150ms, and the final hit navigates after 300ms.
- **No JS:** the middle box is a plain link to `/random`.

**The Parody Stamp Rule.** The lid stamp is a parody folded-paper bird in worn black ink only: no circle, no text, no colour, and never the Snapcraft mark or anything that reads as it.

### Mallet (signature)
Only for a mouse or trackpad (hover and fine pointer); touch devices keep the system pointer.
- **Plate:** a flat vector wooden mallet, tan wood with a black outline in the box's style. Source plate `assets/plates/mallet.png`, generation prompt embedded; the shipped derivatives are a 360 px webp (with a `.json` provenance note) and 128/256 px cursor PNGs (provenance as PNG tEXt).
- **Drawn mallet (with JS):** the system cursor is hidden and the mallet is drawn 180px wide, fixed, following the pointer. It pivots at the end of its handle. At rest it is rotated -30deg, so the head hovers about 55px above the pointer.
- **Strike:** a press swings it down in 170ms: -30deg, overshoot to -58deg on an ease-in, back to -49deg, settling at -52deg, where the striking face sits exactly on the pointer. It stays down at least 180ms; on release it lifts back with a 240ms springy rotate (cubic-bezier(0.2, 1.5, 0.4, 1)).
- **Hidden:** when the pointer leaves the window, over the night band, which shows the normal cursor (the hand on its links), and near the theme picker: within 120px of its centre the page gets `near-ui`, the mallet hides and the cursor is the normal one (the pointer on the button), because the 180px mallet's head is drawn above the pointer and would cover it. A press over the band or near the picker does not swing or play a sound.
- **Without JS:** a 128/256 px mallet cursor image, hotspot on the striking face (28, 59), falling back to the pointer.
- **Reduced motion:** no strike animation and no lift transition; the pose still changes between rest and down.

### Sound
Synthesised with Web Audio, no files, and only ever in answer to a click, tap or key.
- **Hit:** a triangle "bonk" starting at 260 Hz minus 30 Hz per hit landed, this one included (230 Hz, then 200 Hz), and falling to 60 Hz over 0.16s, plus a 0.09s band-passed (1100 Hz) noise crunch, at gain 0.4.
- **Kill:** the final hit, a heavier bonk falling 170 to 42 Hz over 0.24s with a 0.22s crunch, at gain 0.5.
- **Miss:** a mouse press on the board that hits nothing gets a duller thud, 130 to 45 Hz with 0.06s of low-passed noise, at gain 0.25.

### Masthead
- **Style:** "I fucking hate snaps." centred at the top of reading pages, masthead type in ink, no underline; on hover a red underline appears (1.5px, offset 0.22em, no skip-ink), matching the other links. Links home.
- **Focus:** 3px ink outline at 4px offset with 2px corners.

### Pull and Quote (signature, reading pages)
- **Pull:** the first highlight of the quote as the page's h1, in curly quotes, pull type in ink. A quote with no highlight gets a visually hidden "Opinion No. N" h1 and a solo blockquote instead.
- **Quote:** the whole complaint as a real blockquote in soft ink at reading size, paragraphs 0.6em apart.
- **Highlight:** a `mark` with no background: ink, weight 800, a 0.12em ink underline at 0.16em offset. Never colour alone.
- **Meta:** "Opinion No. N · " plus the permanent link in muted ink; the link selects as one unit.
- **404:** "No opinion here." as the pull, "Plenty of hate elsewhere." as a soft-ink lede, and "Random opinion" as an ink pill, because a lone link with no resting underline would read as plain text.

### Ink Pill
The one filled control, for a page's main action: the share button on a quote page and the "Random opinion" link on the 404. Shared by `button` and `a` (inline-block, no text decoration).
- **Shape:** fully rounded pill (999px), no border.
- **Fill:** ink with paper text, button type, padding 0.85em 1.5em.
- **Hover / Focus:** fill turns soft ink; focus is a 3px ink outline at 3px offset.
- **Share behaviour:** "Copy link" copies the permanent URL and says "Copied" for 1.8s (or selects the visible link and says "Link selected" without clipboard access). On a coarse pointer with Web Share it reads "Share" and opens the share sheet. A visually hidden status line repeats the result. Without JS it stays hidden and the link is shown as selectable text.

### Action Links
- **Style:** on a quote page, "Add your own rant" (to the support URL) and "Another opinion" (to `/random`) beside the pill. Ink, no underline at rest: standing apart from text, they read as links by position.
- **Hover:** a red underline, 1.5px, offset 0.22em, no skip-ink.
- **Focus:** 3px ink outline at 4px offset with 2px corners.

### Theme Picker
A small round icon button fixed at the top right of every page (0.75rem in), above everything else. Hidden without JS; the page then follows the system.
- **Shape:** 44px circle, transparent, no border; a 22px inline SVG icon in ink.
- **Cycle:** system, light, dark. The icon shows the current choice: a half-filled circle (system), a ring (light), a crescent (dark). No rays.
- **Label:** "Colour theme: follows your system. Change it.", "Colour theme: light. Change it.", "Colour theme: dark. Change it."
- **Hover / Focus:** an 8% ink wash; focus is a 3px ink outline at 2px offset.
- **Persistence:** a light or dark pick is saved in local storage (`theme`); choosing system removes it. A tiny inline script in the head applies a saved pick before first paint, and the page declares `color-scheme: light dark`.

### Link Preview
Every page carries Open Graph and Twitter large-image tags: title (a quote's punchline in curly quotes, or "Opinion No. N"), description (the whole quote), canonical URL, and `og.jpg`, a 1200 x 630 JPEG cut from the approved landing comp (its origin is in a JPEG comment segment), with alt text describing the box in its hole.

### Report Band
One sentence, after ifuckinghatejira.com: "Real opinions from real people about Ubuntu snaps. Add your own rant, or buy me a coffee and help me survive the next auto-refresh." Both links go to the support URL. It is a real `footer` (the contentinfo landmark), outside `main`, on every page.
- **Band:** Night, with Night Ink text, full width on the bottom edge, no rule, square corners, padding 1.15em clamp(1rem, 4vw, 3.5rem) 1.25em, left-aligned body type. The cursor over it is the normal one.
- **Links:** body type in Night Ink with a hairline underline at rest (1px, Night Ink at 50% via `color-mix`, offset 0.22em, no skip-ink), kept so links inside a sentence stay distinguishable without colour (WCAG 1.4.1).
- **Hover:** the underline turns red at 1.5px.
- **Focus:** 3px Night Ink outline at 4px offset with 2px corners (about 9.7:1 on Night).

## Do's and Don'ts

### Do:
- **Do** keep every glyph in the warm inks on warm paper (charcoal for display and links, soft ink for reading, muted ink for meta), except on the night band, whose text is night ink; keep contrast soft, never pure black type.
- **Do** size the landing page's desktop layout in comp pixels of the 1536 x 1024 design so proportions hold at every size.
- **Do** give every motion a reduced-motion state where movement is instant and the interaction still works.
- **Do** show focus as a 3px outline: ink at 6px offset on the snap's hole, ink at 4px offset on reading-page links (3px on the pill), night ink at 4px offset on the band.
- **Do** mark a quote's highlight with weight and an ink underline, never colour alone.
- **Do** keep the ink pill for a page's one main action (share on a quote page, the way on from the 404).
- **Do** keep "Add your own rant" and the coffee link inside the one closing sentence; a coffee link is allowed there and nowhere else.
- **Do** keep the provenance of every raster plate with its source file in `assets/plates/`, and carry it inside every shipped raster (PNG tEXt, JPEG comment) or, for WebP only, as a `.json` note beside it.

### Don't:
- **Don't** use red for anything except a hovered link's underline.
- **Don't** drop the resting hairline underline from links inside a sentence.
- **Don't** use orange, or any Ubuntu or Snapcraft mark, logo or colour, in CSS or imagery. The cardboard tan inside the box plate is its material, not an accent to reuse.
- **Don't** add gradients, shadows, gloss, riso, grain or texture outside the holes' depth treatment. The worn edge of the lid stamp is part of the box plate, not a surface treatment to reuse.
- **Don't** add rays, bursts or starbursts behind the box or the type.
- **Don't** put quotes on the landing page.
- **Don't** add a third Figtree weight or a second typeface.
