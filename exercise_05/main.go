package main

import (
	"encoding/xml"
	"flag"
	link "github.com/krozlink/gophercises/exercise_04"
	"net/http"
	"strings"
)

func main() {
	depth := flag.Int("depth", -1, "The maximum number of links to follow")
	url := flag.String("url", "", "URL of the target website")

	scrapedLinks := make(map[string]bool)

	result, err := scrape(*url, *depth)
}

func scrape(url string, depth int) ([]link.Link, error) {

	links := make(map[string]bool)

	var scraper func(u string, d int) ([]link.Link, error)
	scraper = func(u string, d int) ([]link.Link, error) {

		pageLinks, err := readLinks(url)
		if err != nil {
			panic(err)
		}

		scrapedLinks[*url] = true

		found := make(map[string]bool)
		for _, l := range links {
			u := l.Href

			if len(l.Href) > 0 && strings.Index(l.Href, "/") == 0 {
				u = *url + l.Href
			}

			if !found[u] {
				found[u] = true
			}
		}

		return nil, nil
	}

	return scraper(url, depth)
}

func readLinks(url string) ([]link.Link, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	links, err := link.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return links, nil
}

type urlset struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []url    `xml:"url"`
}

type url struct {
	Location string `xml:"loc"`
}
