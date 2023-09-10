package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	fileFlag string
)

func init() {
	flag.StringVar(&fileFlag, "f", "problems.csv", "filepath")
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

	totalQuestion := len(records)
	var totalCorrectAnswers int

	fmt.Printf("This quiz contains %d questions. Please answer it and marks will be given by the end of the quiz\n", totalQuestion)

	for i, record := range records {
		correctAnswer, _ := strconv.Atoi(record[1])

		fmt.Printf("Question %d: %s = ", i+1, record[0])
		line, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading solution:", err)
			return
		}

		providedAnswer, err := strconv.Atoi(strings.TrimRight(line, "\n"))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid number:", err)
			return
		}

		if providedAnswer == correctAnswer {
			totalCorrectAnswers++
		}

	}

	fmt.Printf("Total correct answers: %d/%d\n", totalCorrectAnswers, totalQuestion)
}
