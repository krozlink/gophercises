package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var problemFile string
var timer int

func main() {

	problemFile = *flag.String("filename", "problems.csv", "File containing the list of problems")
	timer = *flag.Int("timer", 30, "Number of seconds before running out of time")

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

	timer := time.NewTimer(time.Second * time.Duration(timer))
	answerCh := make(chan string)

loop:
	for _, p := range problems {
		fmt.Println(p.Question)
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()

		select {
		case <-timer.C:
			fmt.Println("Times up!")
			break loop
		case a := <-answerCh:
			if a == p.Answer {
				correct++
			} else {
				incorrect++
			}
		}
	}

	fmt.Printf("You got %v out of %v questions correct\n", correct, count)
}

func readProblems() ([]*Problem, error) {
	file, err := os.Open(problemFile)
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
