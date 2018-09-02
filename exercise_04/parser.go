package parser

import (
	"golang.org/x/net/html"
	"io"
)

// ParseLinks returns a slice of Links parsed from the supplied Reader
func ParseLinks(r io.Reader) []Link {
	parent, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var result []Link
	var parser func(*html.Node)

	parser = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			l := Link{
				Text: parseText(n),
			}

			for _, a := range n.Attr {
				if a.Key == "href" {
					l.Href = a.Val
				}
			}
			result = append(result, l)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parser(c)
		}
	}

	parser(parent)
	return result
}

func parseText(n *html.Node) string {
	var text string
	if n.Type == html.TextNode {
		text += n.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += parseText(c)
	}

	return text
}

type Link struct {
	Href string
	Text string
}
