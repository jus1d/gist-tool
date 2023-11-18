package main

import (
	"fmt"
)

func main() {
	url, err := createGist("go.mod", "")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("gist url: %s", url)
}
