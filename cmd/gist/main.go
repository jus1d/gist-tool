package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/jus1d/gist-tool/internal/gist"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("Incorrect usage\n")
		return
	}

	path := args[0]

	c := gist.New("")
	url, err := c.Create(path)
	if err != nil {
		panic(err)
	}

	err = clipboard.WriteAll(url)
	if err != nil {
		fmt.Println("Your Gist URL: ", url)
	} else {
		fmt.Println("URL copied to clipboard!")
	}
}
