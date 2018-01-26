#!/usr/bin/env bash

set -e

go run cmd/github_css.go
go generate cmd/github_css.go
go build -o ./dist/generate_blog_by_issues main.go
git checkout internal/files/assets.go
