package main

import (
	"fmt"
	"github.com/jus1d/gist-tool/internal/gist"
)

func main() {
	c := gist.New("")
	url, err := c.Create("./gist/gist.go")
	if err != nil {
		panic(err)
	}

	fmt.Printf("gist url: %s", url)
}
