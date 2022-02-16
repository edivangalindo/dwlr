package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/cavaliergopher/grab/v3"
)

func main() {
	maxGoRoutines := 50

	// Check for stdin input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No urls detected. Hint: cat urls.txt | dwlr")
		os.Exit(1)
	}

	sem := make(chan bool, maxGoRoutines)

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		sem <- true
		go func() {

			defer func() { <-sem }()
			// Read urls from stdin

			url := s.Text()

			hostname, filename, err := extract(url)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			// Download the url
			resp, err := grab.Get("./dwlr/"+hostname+"/"+filename, url)

			if err != nil {
				return
			}

			fmt.Println(resp.Filename)

			if err := s.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
				os.Exit(1)
			}
		}()
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
