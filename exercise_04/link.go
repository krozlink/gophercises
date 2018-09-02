package link

import (
	"golang.org/x/net/html"
	"io"
)

// Parse returns a slice of Links parsed from the supplied Reader
func Parse(r io.Reader) ([]Link, error) {
	parent, err := html.Parse(r)
	if err != nil {
		return nil, err
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
	return result, nil
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

// Link represents a link in an HTML document (<a href="">)
type Link struct {
	Href string
	Text string
}
