package main

import (
	"flag"
	"log"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

func getCommandLine() (string, string) {
	repo := flag.String("repo", "", "where the repo is")
	token := flag.String("t", "", "github personal token<todo>")
	flag.Parse()

	if *repo == "" {
		log.Fatal("must set github repo where issue is")
	}
	if *token != "" {
		// todo
		log.Fatal("cannot use token todoing")

	}
	return *repo, *token
}

func main() {
	repo, _ := getCommandLine()

	internal.Run(repo)
}
