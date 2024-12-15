// Fetch prints the content found at a URL.

// Exercise  1.9:  Modify  fetch to  also  print  the  HTTP  status  code,
// found  in resp.Status.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const PREFIX = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, PREFIX) {
			url = PREFIX + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
        fmt.Println(resp.Status);
		resp.Body.Close()
	}
}
