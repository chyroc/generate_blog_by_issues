package internal

//
//import (
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//)
//
//func TestGetIssue(t *testing.T) {
//	t.Skip()
//	as := assert.New(t)
//	issues, err := getIssuesByPage("chyroc/chyroc.github.io", 1, "")
//	as.Nil(err)
//	as.True(len(issues) > 0)
//	as.Equal("", issues[0])
//}
//
//func TestParseHTML(t *testing.T) {
//	t.Skip()
//	as := assert.New(t)
//	html, err := parseToHTML("`~/.zshrc`", "")
//	as.Nil(err)
//	as.Equal("<p><code>~/.zshrc</code></p>\n", html)
//}
//
//func TestNote(t *testing.T) {
//	t.Skip()
//	as := assert.New(t)
//
//	{
//		n := noteImpl{}
//		a, err := n.analysisNote(&content{Name: "name", Content: "IyDmoIfpopgKCi0gdGltZSAyMDE4LTAxLTI1"})
//		as.Nil(err)
//		as.NotNil(a)
//	}
//
//	{
//		n := noteImpl{
//			Repo:  "Chyroc/notes",
//			Paths: []string{"Language/Go/goquiz.github.io/defer-return.md"},
//		}
//		c, err := n.download("Language/Go/goquiz.github.io/defer-return.md")
//		as.Nil(err)
//		as.NotNil(c)
//
//		a, err := n.analysisNote(c)
//		as.Nil(err)
//		as.NotNil(a)
//	}
//}
//
//func TestBase64(t *testing.T) {
//	assert.Equal(t, "IyDmoIfpopgKCi0gdGltZSAyMDE4LTAxLTI1", encodeBase64("# 标题\n\n- time 2018-01-25"))
//
//	s, err := decodeBase64("IyBkZWZlcuS4jnJldHVybueahOmXrumimO+8iGRlZmVy5LmL5LiA77yJCgpb\nZ28taW50ZXJuYWxzLzAzLjQubWQgYXQgbWFzdGVyIMK3IHRpYW5jYWlhbWFv\n")
//	assert.Nil(t, err)
//	assert.Equal(t, "# defer与return的问题（defer之一）\n\n[go-internals/03.4.md at master · tiancaiamao", string(s))
//
//	d, err := decodeBase64(encodeBase64("s"))
//	assert.Nil(t, err)
//	assert.Equal(t, "s", string(d))
//}
