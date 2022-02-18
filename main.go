/*
	Copyright © 2022 edivangalindo
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	threads := flag.Int("t", 50, "Number of threads to utilise.")

	// Check for stdin input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No urls detected. Hint: cat urls.txt | dwlr")
		os.Exit(1)
	}

	results := make(chan string, *threads)

	go func() {

		// Read urls from stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			url := s.Text()

			hostname, filename, err := extract(url)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			// Download the url
			resp, err := grab.Get("./dwlr/"+hostname+"/"+filename, url)
			if err != nil {
				fmt.Println(err)
			}

			printResult("Downloaded: "+resp.Filename, results)
		}

		close(results)
	}()

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for res := range results {
		fmt.Fprintln(w, res)
	}
}

func printResult(url string, results chan string) {
	result := url
	if result != "" {
		results <- result
	}
}

// Extract the hostname and path from a url
func extract(urlString string) (string, string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", "", err
	}

	// Get last resource asked of url
	// Example: https://www.google.com/test.js => test.js
	f := strings.Split(u.String(), "/")
	l := f[len(f)-1]

	return u.Hostname(), l, nil
}
