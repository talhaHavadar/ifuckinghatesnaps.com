---
version: 1
slug: "internal-site-templates-landing-html"
primary_target: "internal/site/templates/landing.html"
related_targets: []
---

# Landing page (`/`)

Mode: Persuade. First visit to ifuckinghatesnaps.com.

## Scope and job

- Audience: Ubuntu users fed up with snaps, arriving cold from a shared link or curiosity.
- Job: land the joke in a second, then get them to a real opinion. Whacking a snap goes to `/random`, which 302s to a permanent `/<n>/`.
- Named success: the visitor submits their own. A full-width night band on the bottom edge (#222a3a with #d9d3c8 text, small type, no top rule; the user asked for "black with dark blue tones" on 2026-11-27; the page never scrolls, and over the band the cursor is the normal one, not the mallet) carries the tagline and both actions in one sentence, after ifuckinghatejira.com: "Real opinions from real people about Ubuntu snaps. Add your own rant, or buy me a coffee and help me survive the next auto-refresh." Both links go to the support URL (Buy Me a Coffee); the user reversed the earlier no-coffee-link rule.
- No quotes on this page. No navigation. Proof is the gag itself; no stats, testimonials or invented content.
- URL contract change: `/` no longer serves the newest quote. Pinned by a test; the PR calls it out. No `/latest/`.
- Anti-goals: slick or corporate, hostile, busy, Ubuntu- or Snapcraft-branded. No orange, no rays or bursts, no riso or grain.

## Direction contract

THESIS: Snaps keep popping up uninvited; here you finally get to whack one. The page is a whack-a-mole board, not a hero with a button. It refuses the headline-over-subcopy-over-CTA landing.

OWN-WORLD: Flat clean vector on a warm off-white ground. Heavy warm-charcoal type (softened from black on 2026-11-27 at the user's request). Three oval holes with depth: a lit far wall falling into darkness and a ring of dirt round the rim. A tan cardboard box with a strip of tape, a smug half-lidded face and "SNAP" on its side. Black text everywhere; the one red is the thin underline on "Report your own". No gradients, textures, shadows or gloss.

STORY: The visitor reads "I fucking hate snaps.", sees the smug box pop out of a hole, whacks it, and lands on a random real opinion. Those who want to vent follow "Report your own".

FIRST VIEWPORT: Centred column on off-white. Headline across the top at display size. Below it, three holes in a row across about 80% of the width, the box popping from one of them (middle at rest). "Whack a snap." large and bold under the holes. A small underlined "Report your own" link at the bottom. Nothing else. The box is the primary action.

FORM: Whack-a-Snap, clean box ("box-pop"), chosen on the decision page; seed:3cc57884191d4593. Approved decision comp: `.impeccable/mocks/decision/r6-box-pop.png`.

Interaction: plain JS (about 3 KB) and CSS, no wasm. The box pops from random holes with a bouncy stretch, overshoot and squash, then wobbles smugly while it waits; it ducks back with a small hop and a visible slide into the hole. It stays put while nobody chases it; within 0.45–0.75 s of the mallet reaching it, it is back in its hole (the 320 ms duck included), then pops up from another hole; one that comes up under a resting mallet notices it without the mouse moving. Its lid carries a worn black-ink stamp of a parody bird (four chunky triangles, no circle, text or colour; never the Snapcraft mark); its pupils (drawn in code over the plate's empty eye whites) follow the mallet. It takes three hits (the user asked for this on 2026-11-27): the first two flinch it and swap in battered plates (snap-box-dmg1: dent, peeled tape, annoyed mouth; snap-box-dmg2: crushed corner, torn flap, plasters, grimace) and, under the mallet, send it fleeing to another hole; the third flattens it (x eyes) and navigates to `/random`. Its accessible name counts the hits left; keyboard hits do not make it flee. With a mouse or trackpad the pointer is a drawn wooden mallet (180 px) held at the end of its handle: a press brings the head down onto the pointer with an overshoot, and the hit lands on the press, with a board jolt and a synthesised bonk and cardboard crunch (a dull thud on a miss; sound only answers a click, tap or key). Without JS it falls back to a 128 px mallet cursor. The "Report your own" link keeps the hand and hides the mallet. Without JS the box is a plain link to `/random`. One snap is always in the tab order; focusing it brings it up and pauses the game; Enter whacks. Reduced motion keeps the game running but drops the bounce, the wobble and the squash.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance

## Open

- Quote page is shaped separately later; until then whacking lands on today's unstyled quote page.
- Saved alternative concept: `.impeccable/concepts/bird-proud/`.

## Hardening round (2026-11-27, from the first critique, 25/36)

- The closing sentence is a real `<footer>` landmark; its links have a visible night-ink focus ring and hover.
- A hidden status region speaks each hit ("Ouch. 2 hits left." ... "Got it. Opening a random opinion.").
- The snap tires: dodges slow as fatigue builds, and after about four dodges or nine seconds without a hit it stops running and sulks, so a player who keeps missing still gets through. A hit gives it a second wind.
- On touch it hops by itself until it tires. On phones the holes form a triangle so the snap is about twice as big.
- Miss thuds only on the board; the kill holds 650 ms; Back starts a fresh game; the site has a favicon.

## Themes (2026-11-27, at the user's request)

Light and dark, following the system setting by default; a round picker in the top-right corner cycles System, Light and Dark and remembers the choice (applied in the head before paint). Dark is the same world with the lights down (softened at the user's request, who found the first dark "fake" and "jarring"): soft warm charcoal paper #24211e, warm light ink #e2dacd a step below white, a quiet neutral band #1b1c20, a lighter hover red #e07a6f, a dirt ring that nearly disappears, and the drawn plates dimmed slightly (brightness .88, saturate .9). On the landing page the mallet is put away within 120 px of the picker. In dark the holes are shaded differently (the user saw the light-mode shading read as raised mounds on a dark ground): darker than the ground everywhere, darkest under the far rim, a faint light on the near lip, and no shadow on the ground.
