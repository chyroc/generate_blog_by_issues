package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
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
	as := assert.New(t)
	html, err := parseToHTML("`~/.zshrc`")
	as.Nil(err)
	as.Equal("<p><code>~/.zshrc</code></p>\n", html)
}

func TestSaveArticle(t *testing.T) {
	//as := assert.New(t)
	saveArticle([]*issue{
		{
			Title:     "zsh",
			Body:      "`~/.zshrc`",
			CreatedAt: time.Now(),
		},
	})
	//as.Nil(err)
	//as.Equal("<p><code>~/.zshrc</code></p>\n", html)
}
