package internal

import (
	"log"
	"strconv"
)

func formatFileNmae(i issue) string {
	return "articles/" + formatTime(i.CreatedAt) + "-" + strconv.Itoa(i.ID) + ".html"
}

func (g *generateBlog) saveArticle(issues []issue) {
	for k, i := range issues {
		go func(k int, i issue) {
			log.Printf("start fetch %d: %s\n", k, i.Title)

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
