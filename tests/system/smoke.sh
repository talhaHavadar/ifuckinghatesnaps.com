#!/usr/bin/env bash
# Runs the built image the way production will (read-only, no capabilities)
# and checks it over HTTP.
#
# Usage: smoke.sh IMAGE
set -euo pipefail

image=${1:?usage: smoke.sh IMAGE}
port=${SMOKE_PORT:-18080}
base=http://127.0.0.1:$port

container=$(podman run --detach --rm \
    --read-only --cap-drop=all --security-opt=no-new-privileges \
    --publish "127.0.0.1:$port:8080" "$image")
trap 'podman rm --force --time 0 "$container" >/dev/null' EXIT

for _ in $(seq 50); do
    curl --silent --fail "$base/healthz" >/dev/null && break
    sleep 0.2
done

failures=0

# check <description> <command...>
check() {
    local desc=$1
    shift
    if "$@"; then
        echo "ok    $desc"
    else
        echo "FAIL  $desc"
        failures=$((failures + 1))
    fi
}

status() { curl --silent --output /dev/null --write-out '%{http_code}' "$base$1"; }

check "healthz answers 200" test "$(status /healthz)" = 200
check "home page shows a quote" sh -c "curl --silent --fail '$base/' | grep -q '<blockquote'"
check "quote 1 is served" test "$(status /1/)" = 200
check "unknown quote is 404" test "$(status /999999/)" = 404
check "random redirects" test "$(status /random)" = 302

if ((failures)); then
    echo "--- container logs"
    podman logs "$container"
    exit 1
fi
