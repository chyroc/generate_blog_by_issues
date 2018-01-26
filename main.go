package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

var version = "v0.3.0"

func getCommandLine() (string, string, bool, []byte) {
	repo := flag.String("repo", "", "where the repo is")
	token := flag.String("token", "", "github token")
	config := flag.String("config", "", "json config file")
	v := flag.Bool("v", false, "generate_blog_by_issues version")
	async := flag.Bool("async", false, "async from issues")
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

	return *repo, *token, *async, file
}

func main() {
	repo, token, async, config := getCommandLine()
	if async {
		internal.Async(repo, token, config)
		os.Exit(0)
	}
	internal.Run(repo, token, config)
}
