package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	problems, err := readProblems()
	if err != nil {
		log.Fatal(err)
	}

	runQuiz(problems)
}

func runQuiz(problems []*Problem) {
	scanner := bufio.NewScanner(os.Stdin)

	count := len(problems)
	correct := 0
	incorrect := 0
	for _, p := range problems {
		fmt.Println(p.Question)
		scanner.Scan()
		answer := scanner.Text()

		if answer == p.Answer {
			correct++
		} else {
			incorrect++
		}
	}

	fmt.Printf("You got %v out of %v questions correct\n", correct, count)
}

func readProblems() ([]*Problem, error) {
	filename := flag.String("filename", "problems.csv", "File containing the list of problems")

	file, err := os.Open(*filename)
	if err != nil {
		return nil, fmt.Errorf("Unable to read the problems file")
	}

	defer file.Close()

	c := csv.NewReader(file)
	lines, err := c.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse the problems file")
	}

	problems := make([]*Problem, len(lines))

	for i, l := range lines {
		p := &Problem{
			Question: l[0],
			Answer:   l[1],
		}
		problems[i] = p
	}

	return problems, nil
}

type Problem struct {
	Question string
	Answer   string
}
