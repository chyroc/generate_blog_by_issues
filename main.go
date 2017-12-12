package main

import (
	"flag"
	"log"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

func getCommandLine() string {
	repo := flag.String("repo", "", "where the repo is")

	if *repo == "" {
		log.Fatal("must set github repo where issue is")
	}

	return *repo
}

func main() {
	repo := getCommandLine()

	internal.Run(repo)
}
