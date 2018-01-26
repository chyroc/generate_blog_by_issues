package github

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	"github.com/Chyroc/generate_blog_by_issues/internal/common"
	"github.com/Chyroc/generate_blog_by_issues/internal/files"
)

type labels struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type issue struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	Labels    []*labels `json:"labels"`
}

type issueInterface interface {
	getIssuesByPage(page int) ([]issue, error)
	getIssuesPage() (int, error)
	GetAllIssues() ([]files.Article, error)
}

var _ issueInterface = (*Issue)(nil)

// Issue Issue
type Issue struct {
	Repo  string
	Token string
}

var linkReg = regexp.MustCompile(`rel="next"(.*?)page=(\d*)&per_page=(.*?)rel="last"`)

func (i issue) toArticle() files.Article {
	return files.Article{
		ID:        strconv.Itoa(i.ID),
		Title:     i.Title,
		Body:      i.Body,
		CreatedAt: i.CreatedAt,
	}
}

func repoToIssueURL(repo string) string {
	return "https://api.github.com/repos/" + common.LinkToGithubRepoName(repo) + "/issues"
}

func (g Issue) getIssuesByPage(page int) ([]issue, error) {
	repo := repoToIssueURL(g.Repo)
	resp, err := common.Get(repo, g.Token, map[string]string{"state": "open", "page": strconv.Itoa(page), "per_page": "25"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issues []issue
	if err := json.Unmarshal(buf, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func (g Issue) getIssuesPage() (int, error) {
	page := 1

	repo := repoToIssueURL(g.Repo)
	resp, err := common.Get(repo, g.Token, map[string]string{"state": "open", "page": strconv.Itoa(1), "per_page": "25"})
	if err != nil {
		return page, err
	}
	defer resp.Body.Close()

	if link, ok := resp.Header["Link"]; ok && len(link) > 0 && link[0] != "" {
		if pages := linkReg.FindStringSubmatch(link[0]); len(pages) > 1 {
			page2, err := strconv.Atoi(pages[2])
			if err != nil {
				return page, err
			}

			return page2, nil
		}
	}

	return page, nil
}

// GetAllIssues GetAllIssues
func (g Issue) GetAllIssues() ([]files.Article, error) {
	issues := make([]files.Article, 0)

	page, err := g.getIssuesPage()
	if err != nil {
		return nil, err
	}

	for i := 1; i < page+1; i++ {
		newIssues, err := g.getIssuesByPage(i)
		if err != nil {
			return nil, err
		}

		for _, v := range newIssues {
			issues = append(issues, v.toArticle())
		}
	}

	return issues, nil
}
