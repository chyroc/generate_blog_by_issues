package internal

import (
	"fmt"
	"io/ioutil"
)

func convertIssueToList(i issue) string {
	labels := ""
	for _, l := range i.Labels {
		labels += fmt.Sprintf(" [#%s]()", l.Name)
	}
	return fmt.Sprintf("### %s\n- [%s](%s)%s\n", formatTime(i.CreatedAt), i.Title, "http://blog.chyroc.cn/"+formatFileNmae(i.CreatedAt), labels)
}

func (g *generateBlog) saveReadme(issues []issue) {
	readme := "Chyroc Blog\n> created by issue\n"
	for _, i := range issues {
		readme += convertIssueToList(i)
	}

	ioutil.WriteFile("README.md", []byte(readme), 0644)
}
