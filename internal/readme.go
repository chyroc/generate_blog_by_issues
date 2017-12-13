package internal

import (
	"fmt"
	"io/ioutil"
	"log"
)

func convertIssueToList(i issue, config conf) string {
	labels := ""
	for _, l := range i.Labels {
		labels += fmt.Sprintf(" [#%s]()", l.Name)
	}
	return fmt.Sprintf("\n### %s\n- [%s](http://%s)%s\n", formatTime(i.CreatedAt), i.Title, config.Host+"/"+formatFileNmae(i), labels)
}

func convertBlogrollList(bs []blogroll) string {
	blogroll := ""
	for _, b := range bs {
		blogroll += fmt.Sprintf("\n- [%s](%s)", b.Name, b.Url)
	}
	return blogroll
}

func (g *generateBlog) saveReadme(issues []issue) {
	readme := fmt.Sprintf("%s\n> created by issue\n\n", g.config.Name)
	for _, i := range issues {
		readme += convertIssueToList(i, g.config)
	}

	blogroll := convertBlogrollList(g.config.Blogrolls)

	readme, err := parseToReadme(readme, blogroll, g.token, g.config)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("index.html", []byte(readme), 0644)
}
