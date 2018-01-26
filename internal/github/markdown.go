package github

import (
	"encoding/json"

	"github.com/Chyroc/generate_blog_by_issues/internal/common"
)

func MarkdownToHTML(issueBody, token string) (string, error) {
	body, err := json.Marshal(map[string]string{"text": issueBody})
	if err != nil {
		return "", err
	}

	htmlBody, err := common.Post("https://api.github.com/markdown", token, body)
	if err != nil {
		return "", err
	}

	return string(htmlBody), nil
}
