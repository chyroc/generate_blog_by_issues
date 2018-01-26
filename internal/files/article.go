package files

import (
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/Chyroc/generate_blog_by_issues/internal/common"
)

// Article Article
type Article struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func formatTime(t time.Time) string {
	return strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(t.Day())
}

// FormatHTMLFileNmae FormatHTMLFileNmae
func FormatHTMLFileNmae(i Article) string {
	return common.ArticlesDir + "/" + formatTime(i.CreatedAt) + "-" + i.ID + ".html"
}

// FormatMDFileNmae FormatMDFileNmae
func FormatMDFileNmae(i Article) string {
	return common.MarkdownsDir + "/" + formatTime(i.CreatedAt) + "-" + i.ID + ".html"
}

// SaveFile SaveFile
func SaveFile(filename, body string) error {
	return ioutil.WriteFile(filename, []byte(body), 0644)
}

// ResetArticlesDir ResetArticlesDir
func ResetArticlesDir() error {
	err := os.RemoveAll(common.ArticlesDir)
	if err != nil {
		return err
	}

	return os.MkdirAll(common.ArticlesDir, 0700)
}

type articleSlice []Article

func (c articleSlice) Len() int {
	return len(c)
}
func (c articleSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c articleSlice) Less(i, j int) bool {
	return c[j].CreatedAt.Before(c[i].CreatedAt)
}

// GroupArticles GroupArticles
func GroupArticles(articles []Article) [][]Article {
	if len(articles) == 0 {
		log.Fatal("must have one issue")
	}

	sort.Sort(articleSlice(articles))

	var is [][]Article
	var currentIssues []Article
	var currentTime = strconv.Itoa(articles[0].CreatedAt.Year()) + "-" + strconv.Itoa(int(articles[0].CreatedAt.Month()))
	for _, i := range articles {
		t := strconv.Itoa(i.CreatedAt.Year()) + "-" + strconv.Itoa(int(i.CreatedAt.Month()))
		if t == currentTime {
			currentIssues = append(currentIssues, i)
		} else {
			currentTime = t
			is = append(is, currentIssues)
			currentIssues = []Article{i}
		}
	}
	is = append(is, currentIssues)
	return is
}
