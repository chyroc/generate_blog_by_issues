package internal

import "time"

type issue struct {
	Title      string
	Body       string
	CreateTime time.Time
}

func getIssuesPage(repo string) (int, error) {
	return 0, nil
}

func getIssuesByPage(repo string, page int) ([]*issue, error) {
	return make([]*issue, 0), nil
}

func getAllIssues(repo string) ([]*issue, error) {
	issues := make([]*issue, 0)

	page, err := getIssuesPage(repo)
	if err != nil {
		return nil, err
	}

	for i := 0; i < page; i++ {
		newIssues, err := getIssuesByPage(repo, i+1)
		filterIssues(newIssues)
		if err != nil {
			return nil, err
		}
		issues = append(issues, newIssues...)
	}

	return issues, nil
}

func filterIssues(issues []*issue) {
}
