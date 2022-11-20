package search

import (
	"log"
	"sync"
)



var mathchers = make(make[String]Matcher)



func Run(searchTerm string) {

	feed,err := RetrieveFeeds()

	if err != nil{
		log.Fatal(err)
	}

	result := make(chan *Result)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))


	for _,feed := range feeds {

		matcher,exists  := mathchers[feed.Type]
		// 
		if !exists {
			matcher = mathchers["default"]
		}

		go func(mathcer Matcher,feed *Feed) {
			Match(mathcer,feed, searchTerm, results)
			waitGroup.Done()
		}(mathcer, feed)
	}

	go func() {
		waitGroup.Wait()

		close(results)

	}()

	Display(results)

}