package internal

import (
	"io/ioutil"
	"log"
	"os"
)

func createAssets() {
	if err := os.MkdirAll("./assets/css", 0700); err != nil {
		log.Fatal(err)
	}

	filename := "./assets/css/style.css"

	if err := ioutil.WriteFile(filename, MustAsset("../assets/style.css"), 0644); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err = f.WriteString("\n"); err != nil {
		log.Fatal(err)
	}

	if _, err = f.Write(MustAsset("../assets/github_markdown.css")); err != nil {
		log.Fatal(err)
	}
}
