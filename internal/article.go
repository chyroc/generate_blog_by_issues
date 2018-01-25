package internal

import (
	"log"
	"time"
)

type article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func formatFileNmae(i article) string {
	return "articles/" + formatTime(i.CreatedAt) + "-" + i.ID + ".html"
}

func (g *generateBlog) saveArticle(issues []article) {
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
