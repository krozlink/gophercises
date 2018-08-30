package main

import (
	parser "github.com/krozlink/gophercises/exercise_04"
	"strings"
)

func main() {
	r := strings.NewReader(`
	<html>
		<head>
			<title>Test</title>
		</head>
		<body>
			<a href="test">Link text</a>
		</body>
	`)
	parser.ParseLinks(r)
}
