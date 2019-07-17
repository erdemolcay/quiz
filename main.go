package main

import (
	"flag"
	"github.com/erdemolcay/quiz/pkg/qns"
	"log"
	"os"

	"github.com/erdemolcay/quiz/pkg/csv"
)

func main()  {
	filePath := flag.String("csv", "", "A CSV file in the format of question,answer,points")
	flag.Parse()
	if !isFlagPassed("csv") {
		log.Fatalln("Please specify CSV file path via csv flag. eg: -csv=quiz.csv")
	}

	if _, err := os.Stat(*filePath); os.IsNotExist(err) {
		log.Fatalf("File: %s does not exist!\n", *filePath)
	}

	questions := csv.Read(*filePath)

	qns.Intro(questions)

	qns.Score(qns.Ask(questions))
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
