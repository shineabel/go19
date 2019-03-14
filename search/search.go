package search

import "sync"
import "log"

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {

	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)
	var wg sync.WaitGroup

	wg.Add(len(feeds))

	for _, f := range feeds {

		matcher, exist := matchers[f.Type]
		if !exist {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			wg.Done()
		}(matcher, f)

	}
	go func() {
		wg.Wait()
		close(results)
	}()

	Display(results)

}

func Register(feedType string, matcher Matcher) {

	if _, exist := matchers[feedType]; exist {
		log.Fatal(feedType, " matcher exist")
	}
	log.Println("register matcher", matcher, "for feed type:", feedType)
	matchers[feedType] = matcher

}
