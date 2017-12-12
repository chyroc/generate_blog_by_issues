package internal

type issue struct {
}

func getIssuesPage(repo string) int {
	return 0
}

func getIssuesByPage(repo string, page int) []*issue {
	return make([]*issue, 0)
}

func getAllIssues(repo string) []*issue {
	issues := make([]*issue, 0)
	for i, j := 0, getIssuesPage(repo); i < j; i++ {
		issues = append(issues, getIssuesByPage(repo, i+1)...)
	}

	return issues
}

func filterIssues(issues []*issue) {
}
