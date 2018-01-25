#!/usr/bin/env bash

set -e

go run cmd/github_css.go
go generate cmd/github_css.go
go build -o generate_blog_by_issues main.go
git stash save
git stash drop
