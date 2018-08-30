package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	// "io/ioutil"
)

func ParseLinks(r io.Reader) []Link {
	parent, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	var result []Link

	var parser func(*html.Node, int)
	parser = func(n *html.Node, depth int) {
		if n.Type == html.ElementNode && n.Data == "a" {
			fmt.Printf("%v\n", n.Data)
			for _, a := range n.Attr {
				fmt.Printf("  %v = %v\n", a.Key, a.Val)
			}
		}

		if n.Type == html.TextNode {
			fmt.Printf("%v\n", n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parser(c, depth+1)
		}
	}

	parser(parent, 0)

	return result
}

type Link struct {
	Href string
	Text string
}
