package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'qustetion,answer'")
	timeLimit := flag.Int("limit", 1, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}
	problems := parseLines(lines)
	correct := 0
	answerCh := make(chan string)

	go func() {
		for {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}
	}()

problemloop:
	for i, problem := range problems {
		timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
		fmt.Printf("Вопрос %v: %v =\n", i+1, problem.q)
		select {
		case <-timer.C:
			fmt.Println("Не успела!")
			continue problemloop
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
			continue problemloop
		}

	}
	fmt.Printf("Ты набрала %d из %d", correct, len(problems))
}
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
