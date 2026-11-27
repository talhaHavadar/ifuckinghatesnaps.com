image := "localhost/ifuckinghatesnaps:latest"

# List recipes
default:
    @just --list

# Run the server from source on http://localhost:8080
dev:
    go run ./cmd/server

# Run the functional tests with the dev shell's Go
test:
    go test ./...

# Run the functional tests with the Go from the Containerfile, which builds the shipped binary
test-image:
    podman build -f Containerfile --target test .

# Run every lint hook
lint:
    prek run --all-files

# Install the git hooks
hooks:
    prek install

# Check the server binary in the image for known vulnerabilities in dependencies and the Go standard library
vuln: image
    #!/usr/bin/env bash
    set -euo pipefail
    dir=$(mktemp -d)
    trap 'rm -rf "$dir"' EXIT
    container=$(podman create {{image}})
    podman cp "$container:/server" "$dir/server"
    podman rm "$container" >/dev/null
    govulncheck -mode=binary "$dir/server"

# Build the container image
image:
    podman build -f Containerfile -t {{image}} .

# Run the container image on http://localhost:8080
run: image
    podman run --rm --read-only --cap-drop=all -p 127.0.0.1:8080:8080 {{image}}

# Build the image and check it over HTTP
system-test: image
    tests/system/smoke.sh {{image}}

# Check that committed changes since BASE come with tests
check-tests base="origin/main":
    prek run --hook-stage manual tests-accompany-changes --from-ref {{base}} --to-ref HEAD

# Everything CI runs
ci base="origin/main": lint (check-tests base) test-image vuln system-test
