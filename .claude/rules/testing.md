---
paths:
  - "cmd/**"
  - "internal/**"
  - "quotes/*.go"
  - "tests/**"
  - "Containerfile"
  - ".containerignore"
---

# Testing

A green PR must mean nothing a visitor would notice broke, so it can be merged.
Tests describe what the site does, not how it does it: drive it over HTTP and
assert on what a browser would receive.

## Tiers

| Tier       | What runs                                                                         | When                                                                                            |
| ---------- | --------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| Functional | The real handler from `site.New`, served by `httptest`, in `tests/`               | Every PR with the image's Go (`just test-image`), locally with the dev shell's Go (`just test`) |
| System     | The built image run read-only by podman and checked with curl, in `tests/system/` | Every PR, `just system-test`                                                                    |

## When you add or change a feature

1. Cover it with a functional test that requests a URL and asserts on the
   status, headers or HTML. Do not assert on unexported functions or internal
   state.
2. If it changes the Containerfile, the runtime user, ports or flags, cover it
   in the system tier as well.
3. Treat URLs as a contract. Quote URLs are shared and linked from elsewhere. A
   change to a path, a status code or a redirect needs a test that pins the new
   behaviour, and the PR calls it out.
4. A bug fix comes with a regression test at the tier where the bug showed up.

`tests-accompany-changes` fails a PR that changes Go code or templates under
`cmd/`, `internal/` or `quotes/` without touching `tests/*_test.go`, or changes
`Containerfile` or `.containerignore` without touching `tests/system/`. The
tests may land in any commit of the PR. A commit that changes no behaviour
(refactor, rename, comment) carries a trailer with the reason instead, for
example `No-Tests: pure rename, covered by existing HTTP tests`. The trailer
exempts only that commit's own changes. "The test is hard to write" is not a
valid reason. Run `just check-tests` before pushing.

New quote text needs no new test. `tests/quotes_test.go` loads every file in
`quotes/`.

## Quote fixtures

Functional tests build their own quotes with `fstest.MapFS` and pass them to
`site.New`. Assert on fixture quotes, not on the published ones in `quotes/`, so
that editing content never breaks a test.

## Small tests: discouraged, not banned

Every test lives in `tests/`, enforced by the `tests-live-in-tests-dir` hook. Go
lets `tests/` import `internal/` packages, so a test there could call internal
functions. Do that only when HTTP cannot reasonably reach the behaviour, and say
why in a comment on the test. The quote markup is tested over HTTP with a table
of fixture quotes, so it is not an exception.

Do not write tests for constructors, getters, plumbing, or one test per
function.

## Keeping green trustworthy

- A flaky test is a bug: fix it or delete it. No automatic retries.
- Never weaken or remove an assertion to make a PR pass without saying so in the
  PR.
