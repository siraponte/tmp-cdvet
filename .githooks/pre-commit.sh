#!/bin/bash

# Intercept errors and abort with a message
set -e
trap "echo 'Operation aborted'" ERR

# Take all modified go files
FILES=$(git diff --cached --name-only --diff-filter=ACMR -- '*.go' | sed 's| |\\ |g')

# If there are no go files, exit
[ -z "$FILES" ] && exit 0

# Lint
echo "$FILES" | xargs -r golangci-lint run

# Format
echo "$FILES" | xargs -r gofmt -s -w

# Re add modified files
echo "$FILES" | xargs git add

exit 0
