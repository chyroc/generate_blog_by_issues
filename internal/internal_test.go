package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIssue(t *testing.T) {
	t.Skip()
	as := assert.New(t)
	issues, err := getIssuesByPage("chyroc/chyroc.github.io", 1, "")
	as.Nil(err)
	as.True(len(issues) > 0)
	as.Equal("", issues[0])
}

func TestParseHTML(t *testing.T) {
	t.Skip()
	as := assert.New(t)
	html, err := parseToHTML("`~/.zshrc`", "")
	as.Nil(err)
	as.Equal("<p><code>~/.zshrc</code></p>\n", html)
}
