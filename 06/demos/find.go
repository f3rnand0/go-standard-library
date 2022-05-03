package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go language. It's one of my favourites"

	if len(os.Args) > 1 {
		searchTerm := os.Args[1]
		result := strings.HasPrefix(sampleString, searchTerm)

		if result {
			fmt.Printf("The sample string starts with %s!\n", searchTerm)
		} else {
			fmt.Printf("The sample string does NOT start with %s\n", searchTerm)
		}
	} else {
		fmt.Println("You must enter a search term")
	}

}
