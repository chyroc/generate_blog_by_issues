package internal

import "log"

func Run(repo string) {
	issues, err := getAllIssues(repo)
	if err != nil {
		log.Fatal(err)
	}

	saveArticle(issues)
	saveTag(issues)
}
