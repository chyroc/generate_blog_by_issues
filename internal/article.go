package internal

import (
	"log"
	"time"
	"strconv"
)

func formatFileNmae(createdAt time.Time) string {
	return "articles/" + strconv.Itoa(createdAt.Year()) + "-" + strconv.Itoa(int(createdAt.Month())) + "-" + strconv.Itoa(createdAt.Day()) + ".html"
}

func (g *generateBlog) saveArticle(issues []issue) {
	for k := range issues {
		log.Printf("start fetch %s\n", issues[k].Title)
	}

	for _, i := range issues {
		go func(i issue) {
			defer g.wg.Done()

			html, err := parseToHTML(i.Body, g.token)
			if err != nil {
				log.Fatal(err)
			}

			i.CreatedAt.Year()
			if err := saveFile(formatFileNmae(i.CreatedAt), html); err != nil {
				log.Fatal(err)
			}
		}(i)
	}
}
