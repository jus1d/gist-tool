package main

import (
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/jus1d/gist-tool/internal/file"
	"github.com/jus1d/gist-tool/internal/gist"
)

func main() {
	var path string
	var description string

	flag.StringVar(&path, "path", "", "Path to file")
	flag.StringVar(&description, "description", "", "Description to your gist")

	flag.Parse()

	for path == "" || !file.Exists(path) {
		fmt.Print("Enter valid path to some file > ")
		_, _ = fmt.Scan(&path)
	}

	c := gist.New("")
	url, err := c.Create(path, description)
	if err != nil {
		panic(err)
	}

	err = clipboard.WriteAll(url)
	if err != nil {
		fmt.Println("Your Gist URL > ", url)
	} else {
		fmt.Println("URL already copied to your clipboard!")
	}
}
