package internal

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

type labels struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type issue struct {
	ID        int       `json:"id"`
	URL       string    `json:"html_url"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Labels    []*labels `json:"labels"`
	CreatedAt time.Time `json:"created_at"`
}

func formatRepo(repo string) string {
	repo = strings.TrimPrefix(repo, "https://")
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimPrefix(repo, "github.com")

	return "https://api.github.com/repos/" + repo + "/issues"
}

func getIssuesByPage(repo string, page int, token string) ([]issue, error) {
	repo = formatRepo(repo)
	body, err := get(repo, token, map[string]string{"state": "open", "page": strconv.Itoa(page), "per_page": "25"})
	if err != nil {
		return nil, err
	}

	var issues []issue
	if err := json.Unmarshal(body, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func getIssuesPage(repo string) (int, error) {
	// todo
	return 1, nil
}

func (g *generateBlog) getAllIssues() ([]issue, error) {
	issues := make([]issue, 0)

	page, err := getIssuesPage(g.repo)
	if err != nil {
		return nil, err
	}

	for i := 0; i < page; i++ {
		newIssues, err := getIssuesByPage(g.repo, i+1, g.token)
		filterIssues(newIssues)
		if err != nil {
			return nil, err
		}
		issues = append(issues, newIssues...)
	}

	return issues, nil
}

func filterIssues(issues []issue) {
}
