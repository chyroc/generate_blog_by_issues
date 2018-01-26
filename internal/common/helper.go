package common

import (
	"encoding/base64"
	"strings"
)

// LinkToGithubRepoName LinkToGithubRepoName
func LinkToGithubRepoName(repo string) string {
	repo = strings.TrimPrefix(repo, "https://")
	repo = strings.TrimPrefix(repo, "http://")
	repo = strings.TrimPrefix(repo, "github.com")
	return repo
}

// EncodeBase64 EncodeBase64
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// DecodeBase64 DecodeBase64
func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
