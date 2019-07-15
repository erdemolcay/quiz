package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Question struct {
	QnText string
	CorrectAnswer string
	Points int
}

func Read(filePath string) []Question {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	questions := make([]Question, 0)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		qnText := record[0]
		correctAnswer := record[1]
		points, err := strconv.Atoi(record[2])
		if err != nil {
			points = 0
		}

		questions = append(questions, Question{qnText, correctAnswer, points})
	}

	return questions
}
