package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var blogFeed = os.Args[1]
	fmt.Printf("Calculating the health of", blogFeed)
	resp, err := http.Get(blogFeed)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetching information from %s failed: %v\n", blogFeed, err)
		os.Exit(1)
	}

	xmlData, err := ioutil.ReadAll(resp.Body)
	// Why does this close have capitals
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading data from %s failed: %v\n", blogFeed, err)
		os.Exit(1)
	}

	fmt.Printf("%s", xmlData)
}
