# ifuckinghatesnaps.com

Real opinions from real people about Ubuntu snaps.

## Adding a quote

Add `quotes/<number>.txt`, one number higher than the newest quote. Wrap the part to highlight in `**`:

```text
Run lsblk on a fresh Ubuntu install and count the loop devices.

**I dare you.**
```

A blank line starts a new paragraph. A quote's number is its URL (`quotes/0002.txt` is `/2/`), so never renumber or reuse one. Deleting a quote makes its URL return 404.

## Development

```sh
nix develop
just hooks         # install the git hooks once
just dev           # serve from source on http://localhost:8080
just test          # functional tests
just ci            # everything CI runs, including the container smoke test
```

`just --list` shows every recipe. The container recipes need podman on the host.

## Running the container

```sh
just image
podman run --rm --read-only --cap-drop=all -p 8080:8080 localhost/ifuckinghatesnaps:latest
```

The image is distroless and runs as a non-root user. The server listens on port 8080 and answers `GET /healthz` for health checks.
