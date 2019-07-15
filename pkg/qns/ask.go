package qns

import (
	"fmt"
	"github.com/erdemolcay/quiz/pkg/csv"
	"log"
)

type Result struct {
	Correct bool
	Points int
	UserAnswer string
	CorrectAnswer string
}

func Ask(questions []csv.Question) ([]Result, int, int) {
	results := make([]Result, 0)
	correctAnswersNum := 0
	totalScore := 0

	for i, question := range questions  {
		userAnswer := ""
		fmt.Printf("\n%d) %s? ", i+1, question.QnText)
		_, err := fmt.Scan(&userAnswer)
		if err != nil {
			log.Fatal(err.Error())
		}

		if userAnswer == question.CorrectAnswer {
			results = append(results, Result{true, question.Points, userAnswer, question.CorrectAnswer})
			correctAnswersNum++
			totalScore += question.Points
		} else {
			results = append(results, Result{false, 0, userAnswer, question.CorrectAnswer})
		}
	}

	return results, correctAnswersNum, totalScore
}
