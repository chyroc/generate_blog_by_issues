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

	issueImpl issueImpl
	noteImpls []note

	wg *sync.WaitGroup
}

func newBlog(repo, token string, configFile []byte) *generateBlog {
	var config conf
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatal(err)
	}

	var ns []note
	for _, v := range config.Notes {
		ns = append(ns, note{
			Repo:  v.Repo,
			Token: token,
			Paths: v.Paths,
		})
	}
	return &generateBlog{
		repo:   repo,
		token:  token,
		config: config,

		issueImpl: issueImpl{repo: repo, token: token},
		noteImpls: ns,

		wg: new(sync.WaitGroup),
	}
}

// Run fetch issues and generate blog
func Run(repo, token string, configFile []byte) {
	g := newBlog(repo, token, configFile)

	var articles []article

	issueArticles, err := g.issueImpl.getAllIssues()
	if err != nil {
		log.Fatal(err)
	}
	articles = append(articles, issueArticles...)

	for _, n := range g.noteImpls {
		noteArticles, err := n.getAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, noteArticles...)
	}

	g.wg.Add(len(articles))
	g.saveArticle(articles)
	g.saveReadme(articles)

	createAssets()

	g.wg.Wait()
}
