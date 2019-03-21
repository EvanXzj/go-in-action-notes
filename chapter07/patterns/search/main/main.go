// This sample program demonstrates how to implement a pattern for
// concurrent requesting results from different systems and either
// wait for all the results to return or just the first one.
package main

import (
	"log"

	"github.com/EvanXzj/go-in-action-notes/chapter07/patterns/search"
)

func main() {
	//(A all results
	results := search.Submit(
		"golang",
		search.Yahoo,
		search.Google,
		search.Bing,
	)

	for _, result := range results {
		log.Printf("mainA : Results : Info : %+v\n", result)
	}

	//(B) first return result and discard the rest
	results = search.Submit(
		"golang",
		search.Yahoo,
		search.Google,
		search.Bing,
		search.OnlyFirst,
	)

	for _, result := range results {
		log.Printf("mainB : Results : Info : %+v\n", result)
	}
}
