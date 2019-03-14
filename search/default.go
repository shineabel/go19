package search

import "log"

type defaultMatcher struct {
}

func init() {

	var matcher defaultMatcher
	Register("default", matcher)
	log.Println("register default matcher...")
}

func (m defaultMatcher) Search(feed *Feed, searctTerm string) ([]*Result, error) {

	return nil, nil
}
