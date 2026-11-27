#!/usr/bin/env bash
# Fail when a change range modifies product code without touching tests.
#
#   Go code in cmd/, internal/, quotes/   needs a change to tests/*_test.go
#   Containerfile, .containerignore       needs a change under tests/system/
#
# Quote text (quotes/*.txt) needs no test: tests/quotes_test.go loads every
# published quote on each run.
#
# Tests may come from any commit in the range. A commit whose message has a
# `No-Tests: <reason>` trailer exempts that commit's own changes only.
#
# Usage: tests-accompany-changes.sh [FROM_REF [TO_REF]]
# Without arguments it reads PRE_COMMIT_FROM_REF / PRE_COMMIT_TO_REF, which
# prek sets for `prek run --from-ref X --to-ref Y`. With no range it skips.
set -euo pipefail

from=${1:-${PRE_COMMIT_FROM_REF:-}}
to=${2:-${PRE_COMMIT_TO_REF:-HEAD}}

if [[ -z $from ]]; then
    echo "tests-accompany-changes: no ref range given, skipping"
    exit 0
fi

code_re='^(cmd|internal|quotes)/.*\.(go|html)$'
code_tests_re='^tests/[^/]+_test\.go$'
image_re='^(Containerfile|\.containerignore)$'
image_tests_re='^tests/system/'

code_tests=0
image_tests=0
uncovered_code=()
uncovered_image=()

while read -r commit; do
    # --no-show-signature: with log.showSignature set, gpg output would land in
    # $reason and make every signed commit look exempt.
    reason=$(git log -1 --no-show-signature --format='%(trailers:key=No-Tests,valueonly)' "$commit")
    exempt=0
    [[ -n ${reason//[[:space:]]/} ]] && exempt=1
    short=$(git rev-parse --short "$commit")

    while read -r file; do
        [[ -z $file ]] && continue
        if [[ $file =~ $code_tests_re ]]; then
            code_tests=1
        elif [[ $file =~ $image_tests_re ]]; then
            image_tests=1
        elif ((exempt)); then
            continue
        elif [[ $file =~ $code_re ]]; then
            uncovered_code+=("$short $file")
        elif [[ $file =~ $image_re ]]; then
            uncovered_image+=("$short $file")
        fi
    done < <(git diff-tree --no-commit-id --name-only -r "$commit")
done < <(git rev-list --no-merges "$from..$to")

status=0

if ((${#uncovered_code[@]})) && ((!code_tests)); then
    status=1
    echo "Code changed without a functional test:"
    printf '  %s\n' "${uncovered_code[@]}" | sort -u
    echo "Add or update a test in tests/."
    echo
fi

if ((${#uncovered_image[@]})) && ((!image_tests)); then
    status=1
    echo "Image build changed without a system test:"
    printf '  %s\n' "${uncovered_image[@]}" | sort -u
    echo "Add or update a check in tests/system/."
    echo
fi

if ((status)); then
    echo "If a commit above changes no behaviour (refactor, rename, comment),"
    echo "add a trailer with the reason to that commit's message instead:"
    echo "  No-Tests: <why no test is needed>"
    echo "See .claude/rules/testing.md."
fi

exit "$status"
