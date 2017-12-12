package internal

import "log"

// Run fetch issue and generate blog
func Run(repo string) {
	issues, err := getAllIssues(repo)
	if err != nil {
		log.Fatal(err)
	}

	saveArticle(issues)
	saveTag(issues)
}
