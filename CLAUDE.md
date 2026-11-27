# ifuckinghatesnaps.com

A site of anonymous quotes about Ubuntu snaps, with part of each quote
highlighted. One Go binary serves every page. The quotes are compiled into it
from `quotes/`. Submissions and donations go through
https://buymeacoffee.com/uphoenix.

Enter the dev shell with `nix develop`. Run `just lint` before you finish a
change and `just ci` before you push. Podman comes from the host, not from the
dev shell.

The testing rules are in `.claude/rules/testing.md`. Before you write a commit
message, PR description, doc or code comment, load the `unslop` skill and follow
it.

## Layout

| Path            | Contents                                                                 |
| --------------- | ------------------------------------------------------------------------ |
| `cmd/server`    | `main`: flags, HTTP server, graceful shutdown                            |
| `internal/site` | Quote parsing and the HTTP handler. `site.New` is the public entry point |
| `quotes/`       | One `<number>.txt` per quote. `**text**` marks the highlight             |
| `tests/`        | Functional tests over HTTP                                               |
| `tests/system/` | Container smoke test, run by `just system-test`                          |

## Rules and what enforces them

Add a rule here only together with its enforcement. If a rule has nothing
enforcing it, fix that before adding more text.

| Rule                                                                            | Enforced by                                                                 |
| ------------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| Formatting, vet, staticcheck, errcheck                                          | `golangci-lint` hook (`.golangci.yml`)                                      |
| `go.mod` is tidy                                                                | `go-mod-tidy` hook                                                          |
| Tests live in `tests/`, not next to the code                                    | `tests-live-in-tests-dir` hook (`prek.toml`)                                |
| Code changes in a PR come with tests                                            | `tests-accompany-changes` (`just check-tests`, part of `just ci`)           |
| Every published quote parses                                                    | `tests/quotes_test.go`, and the server refuses to start                     |
| The shipped binary has no known vulnerability in its dependencies or Go version | `govulncheck` on the binary from the image (`just vuln`, part of `just ci`) |
| Shell scripts, Containerfile and workflows pass their linters                   | `shellcheck`, `hadolint`, `actionlint` hooks                                |
