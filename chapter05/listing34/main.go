package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example <url>")
		os.Exit(-1)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
	defer resp.Body.Close()
}
