//go:generate go-bindata -pkg files -o ../files/assets.go ../../assets/github_markdown.css ../../assets/style.css
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var cssSourceURL = "https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/%s/github-markdown.min.css"
var cssVersionURL = "https://api.cdnjs.com/libraries/github-markdown-css"

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func getLastestVersion() (string, error) {
	body, err := get(cssVersionURL)
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

	body, err := get(fmt.Sprintf(cssSourceURL, version))
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
