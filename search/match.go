package search

import "log"

type Result struct {
	Filed   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searctTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	seatchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, sr := range seatchResults {
		results <- sr
	}

}

func Display(results chan *Result) {

	for result := range results {
		log.Println(result.Filed, result.Content)
	}
}
