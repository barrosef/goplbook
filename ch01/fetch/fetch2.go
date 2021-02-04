//Get a http url and print his body
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	const (
		httpPrefix = "http://"
	)

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:%v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
