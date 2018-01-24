package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

var version = "v0.2.0"

func getCommandLine() (string, string, []byte) {
	repo := flag.String("repo", "", "where the repo is")
	token := flag.String("token", "", "github token")
	config := flag.String("config", "", "json config file")
	v := flag.Bool("v", false, "generate_blog_by_issues version")
	flag.Parse()

	if *v {
		fmt.Printf("generate_blog_by_issues %s\n", version)
		os.Exit(0)
	}

	if *repo == "" {
		log.Fatal("must set github repo where issue is")
	}
	if *token == "" {
		log.Fatal("must set github token")
	}
	if *config == "" {
		log.Fatal("must set json config file")
	}
	file, err := ioutil.ReadFile(*config)
	if err != nil {
		log.Fatal(err)
	}

	return *repo, *token, file
}

func main() {
	repo, token, config := getCommandLine()
	internal.Run(repo, token, config)
}
