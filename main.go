package main

import (
	"flag"
	"log"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

func getCommandLine() (string, string) {
	repo := flag.String("repo", "", "where the repo is")
	token := flag.String("token", "", "github token")
	flag.Parse()

	if *repo == "" {
		log.Fatal("must set github repo where issue is")
	}
	if *token == "" {
		log.Fatal("must set github token")
	}

	return *repo, *token
}

func main() {
	repo, token := getCommandLine()
	internal.Run(repo, token)
}
