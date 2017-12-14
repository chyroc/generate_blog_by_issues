package internal

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
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

var linkReg = regexp.MustCompile(`rel="next"(.*?)page=(\d*)&per_page=(.*?)rel="last"`)

func formatRepo(repo string) string {
	repo = strings.TrimPrefix(repo, "https://")
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimPrefix(repo, "github.com")

	return "https://api.github.com/repos/" + repo + "/issues"
}

func getIssuesByPage(repo string, page int, token string) ([]issue, error) {
	repo = formatRepo(repo)
	resp, err := get(repo, token, map[string]string{"state": "open", "page": strconv.Itoa(page), "per_page": "25"})
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

func getIssuesPage(repo, token string) (int, error) {
	page := 1

	repo = formatRepo(repo)
	resp, err := get(repo, token, map[string]string{"state": "open", "page": strconv.Itoa(1), "per_page": "25"})
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

func (g *generateBlog) getAllIssues() ([]issue, error) {
	issues := make([]issue, 0)

	page, err := getIssuesPage(g.repo, g.token)
	if err != nil {
		return nil, err
	}

	for i := 1; i < page+1; i++ {
		newIssues, err := getIssuesByPage(g.repo, i, g.token)
		if err != nil {
			return nil, err
		}
		issues = append(issues, newIssues...)
	}

	return issues, nil
}
