// Fetch prints the content found at a URL.

// Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst.
// Use it instead of ioutil.ReadAll
// to copy the response body to os.Stdout without requiring a buffer
// large enough to hold the entire stream. Be sure to check the error result of io.Copy.

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
	}
}
