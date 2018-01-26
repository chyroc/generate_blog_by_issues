#!/usr/bin/env bash

set -e

go run internal/cmd/github_css.go
go generate internal/cmd/github_css.go
go build -o ./dist/generate_blog_by_issues main.go
git checkout internal/files/assets.go
