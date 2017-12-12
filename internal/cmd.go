package internal

import (
	"log"
	"sync"
)

type generateBlog struct {
	repo  string
	token string
	wg    *sync.WaitGroup
}

func newBlog(repo, token string) *generateBlog {
	return &generateBlog{
		repo:  repo,
		token: token,
		wg:    new(sync.WaitGroup),
	}
}

// Run fetch issues and generate blog
func Run(repo, token string) {
	g := newBlog(repo, token)

	issues, err := g.getAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	g.wg.Add(len(issues))
	g.saveArticle(issues)
	g.saveTag(issues)

	g.wg.Wait()
}
