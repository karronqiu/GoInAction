package search

import (
	"fmt"
	"log"
)

// Result ...
type Result struct {
	Field   string
	Content string
}

//Matcher ..
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

//Match ..
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatalln(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

//Display ..
func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
