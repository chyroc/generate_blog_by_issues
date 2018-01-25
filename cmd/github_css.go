//go:generate go-bindata -pkg internal -o ../internal/assets.go ../assets/github_markdown.css ../assets/style.css
package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"

	"github.com/Chyroc/generate_blog_by_issues/internal"
)

var cssSourceURL = "https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/%s/github-markdown.min.css"
var cssVersionURL = "https://api.cdnjs.com/libraries/github-markdown-css"

func getLastestVersion() (string, error) {
	resp, err := internal.Get(cssVersionURL, "", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var r struct {
		Version string
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return "", err
	}

	return r.Version, nil
}

func downCSS() error {
	version, err := getLastestVersion()
	if err != nil {
		return err
	}

	resp, err := internal.Get(fmt.Sprintf(cssSourceURL, version), "", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("./assets/github_markdown.css", body, 0644)
}

func main() {
	if err := downCSS(); err != nil {
		panic(err)
	}
}
