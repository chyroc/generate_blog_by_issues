package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
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

type articleSlice []article

func (c articleSlice) Len() int {
	return len(c)
}
func (c articleSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c articleSlice) Less(i, j int) bool {
	return c[j].CreatedAt.Before(c[i].CreatedAt)
}

func groupArticles(articles []article) [][]article {
	if len(articles) == 0 {
		log.Fatal("must have one issue")
	}

	sort.Sort(articleSlice(articles))

	var is [][]article
	var currentIssues []article
	var currentTime = strconv.Itoa(articles[0].CreatedAt.Year()) + "-" + strconv.Itoa(int(articles[0].CreatedAt.Month()))
	for _, i := range articles {
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

func (g *generateBlog) saveReadme(articles []article) {
	readme := fmt.Sprintf("%s\n> created by issue\n\n", g.config.Name)

	groupIss := groupArticles(articles)
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
