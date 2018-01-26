package internal

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/Chyroc/generate_blog_by_issues/internal/files"
	"github.com/Chyroc/generate_blog_by_issues/internal/github"
)

type generateBlog struct {
	repo   string
	token  string
	config conf

	issueImpl github.Issue
	noteImpls []github.Note

	wg *sync.WaitGroup
}

func newBlog(repo, token string, configFile []byte) *generateBlog {
	var config conf
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatal(err)
	}

	var ns []github.Note
	for _, v := range config.Notes {
		ns = append(ns, github.Note{
			Repo:  v.Repo,
			Token: token,
			Paths: v.Paths,
		})
	}
	return &generateBlog{
		repo:   repo,
		token:  token,
		config: config,

		issueImpl: github.Issue{Repo: repo, Token: token},
		noteImpls: ns,

		wg: new(sync.WaitGroup),
	}
}

// Run fetch issues and generate blog
func Run(repo, token string, configFile []byte) {
	g := newBlog(repo, token, configFile)

	var articles []files.Article

	issueArticles, err := g.issueImpl.GetAllIssues()
	if err != nil {
		log.Fatal(err)
	}
	articles = append(articles, issueArticles...)

	for _, n := range g.noteImpls {
		noteArticles, err := n.GetAllNotes()
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, noteArticles...)
	}

	files.CreateAssets()

	g.wg.Add(len(articles))
	g.AsyncToLocalHTML(articles)
	g.saveReadme(articles)
	g.wg.Wait()
}

// Async fetch issues and save files
func Async(repo, token string, configFile []byte) {
	g := newBlog(repo, token, configFile)

	articles, err := g.issueImpl.GetAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	g.wg.Add(len(articles))
	g.AsyncToGithubIssue(articles)
	g.wg.Wait()
}
