package internal

import "log"

// Run fetch issues and generate blog
func Run(repo string) {
	issues, err := getAllIssues(repo)
	if err != nil {
		log.Fatal(err)
	}

	saveArticle(issues)
	saveTag(issues)
}
