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

func repoToContentURL(repo, path string) string {
	return "https://api.github.com/repos/" + linkToGithubRepoName(repo) + "/contents/" + path
}

func (n note) download(path string) (*content, error) {
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

func (n note) analysisNote(c *content) (*article, error) {
	note, err := decodeBase64(c.Content)
	if err != nil {
		return nil, err
	}

	ts := strings.Split(string(note), "\n")
	var createTime time.Time
	for _, v := range ts {
		if strings.HasPrefix(v, "- time ") {
			t, err := time.Parse(time.RFC3339, string(strings.TrimLeft(v, "- time "))+"T15:04:05Z")
			if err != nil {
				return nil, err
			}
			createTime = t
			break
		}
	}

	return &article{
		ID:        encodeBase64(c.Name),
		Title:     c.Name,
		Body:      string(note),
		CreatedAt: createTime,
	}, nil
}

func (n note) getAllNotes() ([]article, error) {
	var s sync.WaitGroup
	var errs = make([]error, len(n.Paths))
	var articles = make([]article, len(n.Paths))

	for k, v := range n.Paths {
		s.Add(1)
		go func(i int) {
			defer s.Done()

			c, err := n.download(v)
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
