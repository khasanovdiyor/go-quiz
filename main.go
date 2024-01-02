package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFlag := flag.CommandLine.String("csv", "problems.csv", "a csv file name (default is 'problems.csv')")
	limitFlag := flag.CommandLine.Int("limit", 0, "the time limit for the quiz in seconds (default is 30)")
	shuffle := false
	flag.CommandLine.BoolVar(&shuffle, "shuffle", false, "if you provide this flag questions will be shuffled")

	flag.Parse()

	if flag.Lookup("h") != nil {
		flag.Usage()
	}

	file, err := os.Open(*csvFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	questionCount := len(lines)

	if shuffle {
		rand.Shuffle(questionCount, func(i, j int) {
			lines[i], lines[j] = lines[j], lines[i]
		})
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	correctAns := 0
	timer := time.AfterFunc(time.Second*time.Duration(*limitFlag), func() {
		logResult(correctAns, questionCount)
		os.Exit(0)
	})
	defer timer.Stop()

	for i, line := range lines {
		question, ans := line[0], strings.ToLower(line[1])
		fmt.Printf("Question #%d: %s: ", i+1, question)

		var userAns string
		fmt.Scanln(&userAns)
		userAns = strings.ToLower(strings.TrimSpace(userAns))

		if strings.EqualFold(userAns, ans) {
			correctAns++
		}
	}

	logResult(correctAns, questionCount)
}

func logResult(ansCount, total int) {
	fmt.Printf("\nYou have scored %d out of %d", ansCount, total)
}
