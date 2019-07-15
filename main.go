package main

import (
	"github.com/erdemolcay/quiz/pkg/qns"
	"log"
	"os"

	"github.com/erdemolcay/quiz/pkg/csv"
)

func main()  {
	if len(os.Args) < 2 {
		log.Fatalln("Please specify CSV file path as the first argument!")
	}

	filePath := os.Args[1]

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Fatalf("File: %s does not exist!\n", filePath)
	}

	questions := csv.Read(filePath)

	qns.Intro(questions)

	qns.Score(qns.Ask(questions))
}

