# Saved concept: "Bird, proud" (landing page)

Saved on 2026-11-27 from the landing-page shape round. The user plans to build on
it later. It was not picked for the current round.

## Concept

Whack-a-Snap. Three black holes in a row. A parody folded-paper bird stands tall
out of the middle hole: chest out, beak up, eyes closed, a smug little smile.
"Whack a snap." sits under the holes. Whacking it plays the squashed state (see
`r6-bird-whacked.png`) and then goes to a random quote. The landing page shows
no quotes.

- Page: heavy black headline "I fucking hate snaps.", the holes, "Whack a snap.",
  and a small "Report your own" link (goes to Buy Me a Coffee). No coffee link,
  no navigation.
- Look: flat, clean vector. No riso, grain, halftone or texture. Warm off-white
  ground, black type and holes, a dull grey-green bird with a black beak. Red
  appears only as a thin underline on the link.
- Parody constraint: the bird must stay clearly unlike the real Snapcraft mark.
  It has fewer, chunkier triangles, a duller grey-green, a black beak (never the
  orange-red one) and a face. Never use the Snapcraft name or wordmark. This
  bends the earlier "not Ubuntu-branded" rule, by the user's choice.
- Build: plain JS (under ~2 KB) plus CSS animation, no wasm. Without JS each
  snap is a link to `/random`. The snaps are keyboard-focusable. With reduced
  motion, the snap stays up instead of popping.

## Files

- `r7-bird-proud.png`: decision comp (codex image generation). The prompt is
  embedded in the PNG and in `r7-bird-proud.png.json`.
- `r6-bird-whacked.png`: the same bird in the hit state, for the squash.
