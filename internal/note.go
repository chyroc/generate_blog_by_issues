package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type note struct {
	Repo  string   `json:"repo"`
	Paths []string `json:"paths"`

	Token string `json:"-"`
}
type content struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Message string `json:"message"`
}

type noteImpl struct {
	Repo  string   `json:"repo"`
	Paths []string `json:"paths"`

	Token string `json:"-"`
}

func repoToContentURL(repo, path string) string {
	return "https://api.github.com/repos/" + linkToGithubRepoName(repo) + "/contents/" + path
}

func (n noteImpl) download(path string) (*content, error) {
	url := repoToContentURL(n.Repo, path)
	resp, err := get(url, n.Token, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var r content
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.Message != "" {
		return nil, fmt.Errorf(r.Message)
	}

	return &r, nil
}

func (n noteImpl) analysisNote(c *content) (*article, error) {
	article := &article{
		ID: encodeBase64(c.Name),
	}

	note, err := decodeBase64(c.Content)
	if err != nil {
		return nil, err
	}

	var newArticleSlice []string

	ts := strings.Split(string(note), "\n")
	for k, v := range ts {
		if k == 0 {
			article.Title = v[2:]
			newArticleSlice = append(newArticleSlice, v)
			continue
		}
		if strings.HasPrefix(v, "- time ") {
			article.CreatedAt, err = time.Parse(time.RFC3339, strings.TrimLeft(v, " -time")+"T15:04:05Z")
			if err != nil {
				return nil, err
			}
			continue
		}

		newArticleSlice = append(newArticleSlice, v)
	}
	article.Body = strings.Join(newArticleSlice, "\n")

	return article, nil
}

func (n noteImpl) getAllNotes() ([]article, error) {
	var s sync.WaitGroup
	var errs = make([]error, len(n.Paths))
	var articles = make([]article, len(n.Paths))

	for k := range n.Paths {
		s.Add(1)
		go func(i int) {
			defer s.Done()

			c, err := n.download(n.Paths[i])
			if err != nil {
				errs[i] = err
				return
			}

			a, err := n.analysisNote(c)
			if err != nil {
				errs[i] = err
				return
			}

			articles[i] = *a
		}(k)
	}

	s.Wait()

	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return articles, nil
}
