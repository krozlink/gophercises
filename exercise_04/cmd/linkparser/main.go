package main

import (
	"fmt"
	link "github.com/krozlink/gophercises/exercise_04"
	"strings"
)

func main() {
	r := strings.NewReader(`
	<html>
	<body>
	  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
	</body>
	</html>
	`)
	links, _ := link.Parse(r)

	for _, l := range links {
		fmt.Printf("Link: %v\nText: %v\n\n", l.Href, l.Text)
	}
}
