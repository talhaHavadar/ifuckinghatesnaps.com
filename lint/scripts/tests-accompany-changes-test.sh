#!/usr/bin/env bash
# Self-test for tests-accompany-changes.sh: builds throwaway repos for each
# scenario and checks the exit status.
set -euo pipefail

# Keep the user's git config (signing, hooks, templates) out of the throwaway repos.
export GIT_CONFIG_GLOBAL=/dev/null GIT_CONFIG_NOSYSTEM=1

script=$(cd "$(dirname "$0")" && pwd)/tests-accompany-changes.sh
work=$(mktemp -d)
trap 'rm -rf "$work"' EXIT
failures=0

# new_repo: fresh repo in $work/<name> with one base commit tagged `base`.
new_repo() {
    repo=$work/$1
    git init -q "$repo"
    git -C "$repo" config user.name test
    git -C "$repo" config user.email test@example.invalid
    git -C "$repo" commit -q --allow-empty -m base
    git -C "$repo" tag base
}

# change <message> <file>...: commit edits to the given files.
change() {
    local msg=$1
    shift
    for f in "$@"; do
        mkdir -p "$repo/$(dirname "$f")"
        echo "$RANDOM" >>"$repo/$f"
        git -C "$repo" add "$f"
    done
    git -C "$repo" commit -q -m "$msg"
}

# expect <pass|fail> <description>
expect() {
    local want=$1 desc=$2 got
    if (cd "$repo" && "$script" base HEAD >/dev/null); then got=pass; else got=fail; fi
    if [[ $got == "$want" ]]; then
        echo "ok    $desc"
    else
        echo "FAIL  $desc (expected $want, got $got)"
        failures=$((failures + 1))
    fi
}

code=internal/site/site.go
template=internal/site/templates/quote.html
code_test=tests/site_test.go
image=Containerfile
image_test=tests/system/smoke.sh

new_repo no-range
if (cd "$repo" && "$script" >/dev/null); then
    echo "ok    skips without a ref range"
else
    echo "FAIL  skips without a ref range"
    failures=$((failures + 1))
fi

new_repo code-no-test
change feature "$code"
expect fail "code change without tests"

new_repo template-no-test
change feature "$template"
expect fail "template change without tests"

new_repo code-same-commit
change feature "$code" "$code_test"
expect pass "code change with tests in the same commit"

new_repo code-later-commit
change feature "$code"
change tests "$code_test"
expect pass "code change with tests in a later commit"

new_repo exempt
change $'refactor\n\nNo-Tests: pure rename' "$code"
expect pass "code change in a No-Tests commit"

new_repo exempt-does-not-leak
change $'refactor\n\nNo-Tests: pure rename' "$code"
change feature cmd/server/main.go
expect fail "No-Tests on one commit does not cover another"

new_repo empty-trailer
change $'refactor\n\nNo-Tests:' "$code"
expect fail "empty No-Tests trailer does not exempt"

new_repo image-no-test
change image "$image"
expect fail "image change without system tests"

new_repo image-with-test
change image "$image" "$image_test"
expect pass "image change with system tests"

new_repo image-functional-test-only
change image "$image" "$code_test"
expect fail "functional tests do not cover image changes"

new_repo system-test-only-for-code
change feature "$code" "$image_test"
expect fail "system tests do not cover code changes"

new_repo quote-text
change quote quotes/0042.txt
expect pass "new quote text needs no test"

new_repo unrelated
change misc README.md Justfile flake.nix go.mod
expect pass "changes outside code and image need no tests"

((failures == 0)) || { echo "$failures scenario(s) failed"; exit 1; }
