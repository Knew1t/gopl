// / Fetchall fetches URLs in parallel and reports their times and sizes.

// Exercise 1.10:
// Modify fetchall to print its output to a file so it can be examined.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
	for range os.Args[1:] {
		f.Write([]byte(<-ch)) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	f.Close()
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
