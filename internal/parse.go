package internal

import (
	"bytes"
	"encoding/json"
	"html"
	"html/template"
	"io/ioutil"
	"strconv"
	"time"
)

type blogroll struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type conf struct {
	Title     string     `json:"title"`
	Name      string     `json:"name"`
	Host      string     `json:"host"`
	Author    string     `json:"author"`
	Blogrolls []blogroll `json:"blogrolls"`
	Blogroll  string     `json:"blogroll"`
	Content   string     `json:"content"`
}

func formatTime(t time.Time) string {
	return strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
}

func parseToHTML(issueBody, token string) (string, error) {
	body, err := json.Marshal(map[string]string{"text": issueBody})
	if err != nil {
		return "", err
	}

	htmlBody, err := post("https://api.github.com/markdown", token, body)
	if err != nil {
		return "", err
	}

	return string(htmlBody), nil
}

func parseToReadme(issueBody, blogroll, token string, config conf) (string, error) {
	config.Title = config.Name

	htmlBody, err := parseToHTML(issueBody, token)
	if err != nil {
		return "", err
	}
	config.Content = htmlBody

	blogrollBody, err := parseToHTML(blogroll, token)
	if err != nil {
		return "", err
	}
	config.Blogroll = blogrollBody

	var doc bytes.Buffer
	t := template.Must(template.New("readmeTmpl").Parse(readmeTmpl))
	if err := t.Execute(&doc, config); err != nil {
		return "", err
	}

	return html.UnescapeString(doc.String()), nil
}

func parseToArticle(title, issueBody, token string, config conf) (string, error) {
	config.Title = title + " - " + config.Name

	htmlBody, err := parseToHTML(issueBody, token)
	if err != nil {
		return "", err
	}
	config.Content = htmlBody

	var doc bytes.Buffer
	t := template.Must(template.New("articleTmpl").Parse(articleTmpl))
	if err := t.Execute(&doc, config); err != nil {
		return "", err
	}

	return html.UnescapeString(doc.String()), nil
}

func saveFile(filename, html string) error {
	return ioutil.WriteFile(filename, []byte(html), 0644)
}
