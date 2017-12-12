package internal

import (
	"log"
	"strconv"
	"time"
)

func formatFileNmae(createdAt time.Time) string {
	return "articles/" + strconv.Itoa(createdAt.Year()) + "-" + strconv.Itoa(int(createdAt.Month())) + "-" + strconv.Itoa(createdAt.Day()) + ".html"
}

func (g *generateBlog) saveArticle(issues []issue) {

	for _, i := range issues {
		go func(i issue) {
			log.Printf("start fetch %s\n", i.Title)

			defer g.wg.Done()

			html, err := parseToHTML(i.Body, g.token)
			if err != nil {
				log.Fatal(err)
			}

			if err := saveFile(formatFileNmae(i.CreatedAt), html); err != nil {
				log.Fatal(err)
			}
		}(i)
	}
}
