package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Score struct {
	question int
	answer   int
}

func (s Score) printScore() {
	fmt.Printf("Total correct answers: %d/%d\n", s.answer, s.question)
}

var (
	fileFlag    string
	timeoutFlag int
	shuffle     bool
)

func init() {
	flag.StringVar(&fileFlag, "f", "test/fixtures/problems.csv", "File containing the quiz")
	flag.IntVar(&timeoutFlag, "t", 0, "Set time limit to the quiz")
	flag.BoolVar(&shuffle, "s", false, "Shuffle the questions")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	path := *&fileFlag

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading csv:", err)
		return
	}

	records = validateQuestions(records)

	if shuffle {
		randomizeSlice(records)
	}

	scoreChan := make(chan Score)
	go RunQuiz(records, scoreChan, timeoutFlag)

	score := <-scoreChan
	score.printScore()
}

func RunQuiz(records [][]string, scoreChan chan<- Score, duration int) {
	totalQuestion := len(records)
	score := Score{
		question: totalQuestion,
	}

	if duration != 0 {
		timeout := time.Duration(duration) * time.Second
		timer := time.NewTimer(timeout)
		defer timer.Stop()

		go func() {
			<-timer.C
			fmt.Println("\nTime's up! Quiz is over.")
			scoreChan <- score
		}()
	}

	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Printf("This quiz contains %d questions. Please answer it and marks will be given by the end of the quiz\n", totalQuestion)
	if duration != 0 {
		fmt.Printf("You have %d seconds to answer all questions\n", totalQuestion)
	}

	for i, record := range records {
		correctAnswer, _ := strconv.Atoi(record[1])

		var (
			line           string
			providedAnswer int
			err            error
		)

		for {
			fmt.Printf("Question %d: %s = ", i+1, record[0])

			if line, err = stdinReader.ReadString('\n'); err != nil {
				fmt.Fprintln(os.Stderr, "Error reading solution:", err)
				return
			}

			if providedAnswer, err = strconv.Atoi(strings.TrimRight(line, "\n")); err != nil {
				fmt.Fprintln(os.Stderr, "Invalid number:", err)
				continue
			}

			break
		}

		if providedAnswer == correctAnswer {
			score.answer++
		}
	}

	scoreChan <- score
}

func validateQuestions(records [][]string) [][]string {
	var result [][]string
	var err error
	for _, record := range records {
		if err = validQuestion(record[0]); err != nil {
			break
		}
		if err = validAnswer(record[1]); err != nil {
			break
		}

		result = append(result, []string{record[0], record[1]})

	}
	return result
}

func validQuestion(q string) error {
	return nil
}

func validAnswer(a string) error {
	_, err := strconv.Atoi(a)
	if err != nil {
		return errors.New("Invalid number")
	}
	return nil
}

func randomizeSlice(slice [][]string) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
