---
version: 1
slug: "internal-site-templates-quote-html"
primary_target: "internal/site/templates/quote.html"
related_targets: ["internal/site/templates/notfound.html","internal/site/templates/layout.html","internal/site/static/site.css"]
---

# Quote page (`/<n>/`) and 404

Mode: Read. Most visitors arrive cold from a shared link; the rest straight after killing the snap on the landing page. Built code-led in the established world (DESIGN.md).

## Scope and job

- Job: get the joke in a second, see the whole complaint, pass it on.
- Primary action (user's choice): share this quote. "Copy link" copies the permanent URL; on touch with the Web Share API it is "Share" and opens the share sheet. Without JS the link is shown as selectable text.
- Secondary (renamed by the user on 2026-11-27): "Add your own rant" (to the support URL; matched to the footer at the user's request) and "Another opinion" (to /random), as plain links with no underline until a thin red one on hover. The previous quote is still linked as rel="prev" in the head. "Add your own rant" stays in the footer band, whose links keep a hairline underline at rest and turn red on hover.
- Link previews per quote: title is the punchline in quotes (or "Opinion No. N"), description the whole quote, canonical URL, og.jpg cut from the approved landing comp.
- 404 in the same frame: "No opinion here." / "Plenty of hate elsewhere." / Random opinion.
- Unchanged: quote URLs, /random, the quote markup tests rely on (blockquote id, mark, rel="prev").

## Direction contract

THESIS: The punchline is the headline. The page leads with the highlighted line pulled out huge, the way people would quote it, then gives the whole complaint at reading size for context. It refuses the centred card with a quote icon.

OWN-WORLD: The landing world: warm paper ground, warm-charcoal Figtree 800 for the punchline, Figtree 400 reading text in soft ink, the punchline inside the quote marked by weight and an ink underline, red only as link underlines, one ink pill button, the night footer band.

STORY: The visitor reads the punchline, glances at the full complaint, sees it is opinion No. N at a permanent link, copies or shares it, and either reads another or leaves.

FIRST VIEWPORT: Small "I fucking hate snaps." masthead centred at the top. Centred column: the punchline in quotes at display size (about 104px at 1440), the full quote below at about 26px in a 34em measure, "Opinion No. N · ifuckinghatesnaps.com/N/" in muted ink, then the Copy link pill beside Previous / Random links. Night footer on the bottom edge; one screen, no scroll.

FORM: Punchline first, position 6 on my ordered list, chosen on the decision page; seed:b4a031d547acb5c3. Chosen sketch: `.impeccable/mocks/decision/q-punchline.html`.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance

## Open

- Per-quote preview images (rendering the punchline into the image) need an image library in the binary; one generic image for now.

## Themes (2026-11-27, at the user's request)

Light and dark, following the system setting by default; a round picker in the top-right corner cycles System, Light and Dark and remembers the choice (applied in the head before paint). Dark is the same world with the lights down (softened at the user's request, who found the first dark "fake" and "jarring"): soft warm charcoal paper #24211e, warm light ink #e2dacd a step below white, a quiet neutral band #1b1c20, a lighter hover red #e07a6f, a dirt ring that nearly disappears, and the drawn plates dimmed slightly (brightness .88, saturate .9). On the landing page the mallet is put away within 120 px of the picker.

