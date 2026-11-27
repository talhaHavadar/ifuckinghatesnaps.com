# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

Ubuntu users who are fed up with snaps: auto-refreshes at the worst moment,
loop devices cluttering `lsblk`, a `~/snap` directory that keeps coming back.
They usually land on a single quote from a link shared in a forum, chat or
social feed, read it, laugh, maybe hit "random" a few times, and pass a link on.
Some of them come back to submit their own complaint.

## Product Purpose

A joke site that is also a pressure valve. It publishes anonymous, real
complaints about Ubuntu snaps, one per page, with the punchline highlighted.
Success is a visitor who laughs, recognises their own pain, shares a quote
link, and possibly submits their own or buys a coffee.

## Positioning

It follows the single-rant-domain format of https://ifuckinghatejira.com/,
which inspired it (headline that is the domain, "real opinions from real
people", submissions and donations through Buy Me a Coffee). It narrows that
format to one complaint per page, each with a permanent numbered URL meant for
sharing, and a highlighted punchline inside each quote. The reference is
inspiration for format and tone, not something to copy.

## Operating Context

- Visitors arrive mostly on a quote URL (`/<n>/`), not the homepage. `/` is a
  whack-a-snap game with no quote on it: whacking the snap goes to `/random`,
  which redirects to a random `/<n>/`. Navigation on a quote page is "Another
  opinion" (the previous quote) and "Random opinion" (`/random`).
- Quote URLs get shared and linked from elsewhere. They are permanent: never
  renumbered or reused. A deleted quote's URL returns 404.
- Submissions and donations both go to https://buymeacoffee.com/uphoenix.
  There is no submission form on the site itself.
- Quotes are curated by hand and added as `quotes/<number>.txt`, compiled into
  the binary.

## Capabilities and Constraints

- One Go binary serves every page from `html/template` templates in
  `internal/site/templates/`. Pages: quote, 404, plus `/healthz`.
- A quote is one or more paragraphs. `**text**` marks the highlighted part.
- Quote pages must unfurl well when shared in chat and social apps (Open Graph
  and similar previews per quote). Not built yet.
- No front-end constraint on JavaScript or third-party assets was set. Open
  decision.

## Brand Commitments

- Name and headline: "I fucking hate snaps." Profanity is part of the brand.
  Never sanitise it, in the chrome or in quotes.
- Tagline: "Real opinions from real people about Ubuntu snaps."
- Voice: blunt, profane, self-deprecating, funny. Non-quote copy is short, e.g.
  "No opinion here. Plenty of hate elsewhere."
- Quotes are anonymous: no names, handles or avatars, ever.
- The anger targets snaps (the technology and the decisions behind it), not
  Ubuntu users or Canonical staff as people.

## Evidence on Hand

- Six published quotes in `quotes/0001.txt` to `quotes/0006.txt`.
- Every quote is a real submission. Never write, invent or paraphrase quotes,
  not even as placeholders or design samples. Use the published ones or
  test fixtures.
- No logo, imagery, testimonials, traffic numbers or press exist. Do not
  fabricate them.

## Product Principles

1. The quote is the product. Everything else on the page serves reading it and
   sharing it.
2. Funny first. It is a joke site, and the joke should land in the first
   second.
3. Real and anonymous. Authenticity is the whole premise, so nothing is made up
   and nobody is named.
4. Links are forever. A quote URL shared today works the same way years from
   now.
5. Punch at snaps, not at people.

## Accessibility & Inclusion

Target WCAG 2.2 AA: colour contrast (including the highlighted text), full
keyboard use, and correct semantics for screen readers (the quote is a real
`blockquote`, and the highlight must not depend on colour alone).
