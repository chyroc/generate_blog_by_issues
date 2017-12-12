package internal

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

func formatTime(t time.Time) string {
	return strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
}

func parseToHTML(issueBody, token string) (string, error) {
	body, err := json.Marshal(map[string]string{"text": issueBody})
	if err != nil {
		return "", err
	}

	html, err := post("https://api.github.com/markdown", token, body)
	if err != nil {
		return "", err
	}

	return string(html), nil
}

func saveFile(filename, html string) error {
	return ioutil.WriteFile(filename, []byte(articleTmplHeader+html+articleTmplFooter), 0644)
}
