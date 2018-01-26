package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/Chyroc/generate_blog_by_issues/internal/common"
	"github.com/Chyroc/generate_blog_by_issues/internal/files"
)

type content struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Message string `json:"message"`
}

type noteInterface interface {
	download(path string) (*content, error)
	analysisNote(c *content) (*files.Article, error)
	GetAllNotes() ([]files.Article, error)
}

var _ noteInterface = (*Note)(nil)

// Note Note
type Note struct {
	Repo  string   `json:"Repo"`
	Paths []string `json:"paths"`

	Token string `json:"-"`
}

func repoToContentURL(repo, path string) string {
	return "https://api.github.com/repos/" + common.LinkToGithubRepoName(repo) + "/contents/" + path
}

func (n Note) download(path string) (*content, error) {
	url := repoToContentURL(n.Repo, path)
	resp, err := common.Get(url, n.Token, nil)
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

func (n Note) analysisNote(c *content) (*files.Article, error) {
	article := &files.Article{
		ID: common.EncodeBase64(c.Name),
	}

	note, err := common.DecodeBase64(c.Content)
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

// GetAllNotes GetAllNotes
func (n Note) GetAllNotes() ([]files.Article, error) {
	var s sync.WaitGroup
	var errs = make([]error, len(n.Paths))
	var articles = make([]files.Article, len(n.Paths))

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
