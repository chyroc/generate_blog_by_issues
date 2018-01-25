package internal

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
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

func (i issue) toArticle() article {
	return article{
		ID:        strconv.Itoa(i.ID),
		Title:     i.Title,
		Body:      i.Body,
		CreatedAt: i.CreatedAt,
	}
}

type issueImpl struct {
	repo  string
	token string
}

var linkReg = regexp.MustCompile(`rel="next"(.*?)page=(\d*)&per_page=(.*?)rel="last"`)

func repoToIssueURL(repo string) string {
	return "https://api.github.com/repos/" + linkToGithubRepoName(repo) + "/issues"
}

func getIssuesByPage(repo string, page int, token string) ([]issue, error) {
	repo = repoToIssueURL(repo)
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

	repo = repoToIssueURL(repo)
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

func (g issueImpl) getAllIssues() ([]article, error) {
	issues := make([]article, 0)

	page, err := getIssuesPage(g.repo, g.token)
	if err != nil {
		return nil, err
	}

	for i := 1; i < page+1; i++ {
		newIssues, err := getIssuesByPage(g.repo, i, g.token)
		if err != nil {
			return nil, err
		}

		for _, v := range newIssues {
			issues = append(issues, v.toArticle())
		}
	}

	return issues, nil
}
