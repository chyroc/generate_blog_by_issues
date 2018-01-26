package internal

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

var articlesDir = "articles"

type article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func formatFileNmae(i article) string {
	return articlesDir + "/" + formatTime(i.CreatedAt) + "-" + i.ID + ".html"
}

func saveFile(filename, html string) error {
	return ioutil.WriteFile(filename, []byte(html), 0644)
}

func resetArticlesDir() error {
	err := os.RemoveAll(articlesDir)
	if err != nil {
		return err
	}

	return os.MkdirAll(articlesDir, 0700)
}

func (g *generateBlog) saveArticle(issues []article) {
	if err := resetArticlesDir(); err != nil {
		log.Fatal(err)
	}

	for k, i := range issues {
		go func(k int, i article) {
			log.Printf("start fetch %d:\t%s\n", k, i.Title)

			defer g.wg.Done()

			html, err := parseToArticle(i.Title, i.Body, g.token, g.config)
			if err != nil {
				log.Fatal(err)
			}

			if err := saveFile(formatFileNmae(i), html); err != nil {
				log.Fatal(err)
			}
		}(k, i)
	}
}
