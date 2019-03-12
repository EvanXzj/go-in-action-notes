package main

import (
	"flag"
	"fmt"
	"github.com/EvanXzj/go-in-action-notes/chapter03/words"
	"io/ioutil"
)

func main() {
	var filename string

	flag.Parse()
	if flag.NArg() >= 1 {
		filename = flag.Arg(0)
	} else {
		fmt.Println("Please input filename you want to count")
		return
	}

	//fmt.Println(filename)
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	count := words.CountWords(string(contents))
	fmt.Printf("There are %d words in your text.\n", count)
}
