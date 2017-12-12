package internal

func parseToHTML(issueBody string) (string, error) {
	return "", nil
}

func saveFile(filename, html string) error {
	return nil
}

func saveIssue(filename, body string) error {
	html, err := parseToHTML(body)
	if err != nil {
		return err
	}

	return saveFile(filename, html)
}
