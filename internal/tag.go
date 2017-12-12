package internal

import "log"

func saveTag(issues []*issue) {
	for _, issue := range issues {
		html, err := parseToHTML(issue.Body)
		if err != nil {
			log.Fatal(err)
		}

		saveFile("", html)
	}
}
