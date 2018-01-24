package internal

import (
	"encoding/json"
	"log"
	"sync"
)

type generateBlog struct {
	repo   string
	token  string
	config conf
	wg     *sync.WaitGroup
}

func newBlog(repo, token string, configFile []byte) *generateBlog {
	var config conf
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatal(err)
	}
	return &generateBlog{
		repo:   repo,
		token:  token,
		wg:     new(sync.WaitGroup),
		config: config,
	}
}

// Run fetch issues and generate blog
func Run(repo, token string, configFile []byte) {
	g := newBlog(repo, token, configFile)

	issues, err := g.getAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	g.wg.Add(len(issues))
	g.saveArticle(issues)
	g.saveReadme(issues)

	g.wg.Wait()
}
