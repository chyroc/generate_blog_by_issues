package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

func getCommandLine() (string, string) {
	repo := flag.String("repo", "", "where the repo is")
	token := flag.String("token", "", "github token")
	v := flag.Bool("v", false, "generate_blog_by_issues version")
	flag.Parse()

	if *v {
		fmt.Printf("generate_blog_by_issues 0.1.0\n")
		os.Exit(0)
	}

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
