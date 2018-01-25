package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func convertIssueToList(i article, config conf) string {
	return fmt.Sprintf("\n- [%s](http://%s)\n", i.Title, config.Host+"/"+formatFileNmae(i))
}

func convertBlogrollList(bs []blogroll) string {
	blogroll := ""
	for _, b := range bs {
		blogroll += fmt.Sprintf("\n- [%s](%s)", b.Name, b.URL)
	}
	return blogroll
}

func groupIssues(issues []article) [][]article {
	if len(issues) == 0 {
		log.Fatal("must have one issue")
	}

	var is [][]article
	var currentIssues []article
	var currentTime = strconv.Itoa(issues[0].CreatedAt.Year()) + "-" + strconv.Itoa(int(issues[0].CreatedAt.Month()))
	for _, i := range issues {
		t := strconv.Itoa(i.CreatedAt.Year()) + "-" + strconv.Itoa(int(i.CreatedAt.Month()))
		if t == currentTime {
			currentIssues = append(currentIssues, i)
		} else {
			currentTime = t
			is = append(is, currentIssues)
			currentIssues = []article{i}
		}
	}
	is = append(is, currentIssues)
	return is
}

func (g *generateBlog) saveReadme(issues []article) {
	readme := fmt.Sprintf("%s\n> created by issue\n\n", g.config.Name)

	groupIss := groupIssues(issues)
	for _, iss := range groupIss {
		readme += fmt.Sprintf("\n### %s", strconv.Itoa(iss[0].CreatedAt.Year())+"-"+strconv.Itoa(int(iss[0].CreatedAt.Month())))
		for _, i := range iss {
			readme += convertIssueToList(i, g.config)
		}
	}

	blogroll := convertBlogrollList(g.config.Blogrolls)

	readme, err := parseToReadme(readme, blogroll, g.token, g.config)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("index.html", []byte(readme), 0644)
}
