package internal

import (
	"bytes"
	"fmt"
	"github.com/Chyroc/generate_blog_by_issues/internal/blog"
	"github.com/Chyroc/generate_blog_by_issues/internal/files"
	"github.com/Chyroc/generate_blog_by_issues/internal/github"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"strconv"
	"os"
	"github.com/Chyroc/generate_blog_by_issues/internal/common"
)

type blogroll struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type note struct {
	Repo  string   `json:"repo"`
	Paths []string `json:"paths"`
}

type conf struct {
	Title     string     `json:"title"`
	Name      string     `json:"name"`
	Host      string     `json:"host"`
	Author    string     `json:"author"`
	Notes     []note     `json:"notes"`
	Blogrolls []blogroll `json:"blogrolls"`
	Blogroll  string     `json:"blogroll"`

	Content string `json:"content"`
}

func convertIssueToList(i files.Article, config conf) string {
	return fmt.Sprintf("\n- [%s](http://%s)\n", i.Title, config.Host+"/"+files.FormatHTMLFileNmae(i))
}

func convertBlogrollList(bs []blogroll) string {
	blogroll := ""
	for _, b := range bs {
		blogroll += fmt.Sprintf("\n- [%s](%s)", b.Name, b.URL)
	}
	return blogroll
}

func parseToReadme(issueBody, blogroll, token string, config conf) (string, error) {
	config.Title = config.Name

	htmlBody, err := github.MarkdownToHTML(issueBody, token)
	if err != nil {
		return "", err
	}
	config.Content = htmlBody

	blogrollBody, err := github.MarkdownToHTML(blogroll, token)
	if err != nil {
		return "", err
	}
	config.Blogroll = blogrollBody

	var doc bytes.Buffer
	t := template.Must(template.New("readmeTmpl").Parse(blog.TmplReadme))
	if err := t.Execute(&doc, config); err != nil {
		return "", err
	}

	return html.UnescapeString(doc.String()), nil
}

func parseToArticle(title, issueBody, token string, config conf) (string, error) {
	config.Title = title + " - " + config.Name

	htmlBody, err := github.MarkdownToHTML(issueBody, token)
	if err != nil {
		return "", err
	}
	config.Content = htmlBody

	var doc bytes.Buffer
	t := template.Must(template.New("articleTmpl").Parse(blog.TmplArticle))
	if err := t.Execute(&doc, config); err != nil {
		return "", err
	}

	return html.UnescapeString(doc.String()), nil
}

func (g generateBlog) saveReadme(articles []files.Article) {
	readme := fmt.Sprintf("%s\n> created by issue\n\n", g.config.Name)

	groupIss := files.GroupArticles(articles)
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

func (g generateBlog) AsyncToLocalHTML(as []files.Article) {
	if err := files.ResetArticlesDir(); err != nil {
		log.Fatal(err)
	}

	for k, i := range as {
		go func(k int, i files.Article) {
			defer g.wg.Done()

			log.Printf("start fetch %d:\t%s\n", k, i.Title)

			html2, err := parseToArticle(i.Title, i.Body, g.token, g.config)
			if err != nil {
				log.Fatal(err)
			}

			if err := files.SaveFile(files.FormatHTMLFileNmae(i), html2); err != nil {
				log.Fatal(err)
			}
		}(k, i)
	}
}

func (g generateBlog) AsyncToLocalMD(as []files.Article) {
	if err := os.MkdirAll(common.MarkdownsDir, 0700); err != nil {
		log.Fatal(err)
	}

	for k, i := range as {
		go func(k int, i files.Article) {
			defer g.wg.Done()

			log.Printf("start fetch %d:\t%s\n", k, i.Title)

			if err := files.SaveFile(files.FormatMDFileNmae(i), i.Body); err != nil {
				log.Fatal(err)
			}
		}(k, i)
	}
}
