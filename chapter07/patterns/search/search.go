package search

import "log"

type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

// Searcher declares an interface used to leverate different
// search engines to find results.
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

func OnlyFirst(s *searchSession) {
	s.first = true
}

func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	// Perfor the searches concurrently.Using a map because
	// it returns the searchers in a random order every time.
	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result

	// Wait for the results to come back
	for search := 0; search < len(session.searchers); search++ {
		// If we just want the first result, don't wait any longer by
		// concurrently discarding the remaining searchResults.
		// Failing to do so will leave the Searchers blocked forever
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		// Wait to recieve results.
		log.Println("search : Submit : Info : Waiting For Results...")
		result := <-session.resultChan

		// Save the results to the final slice.
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		results = append(results, result...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}
