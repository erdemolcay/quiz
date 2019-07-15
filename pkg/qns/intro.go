package qns

import (
	"fmt"
	"github.com/erdemolcay/quiz/pkg/csv"
	"log"
)

func Intro(questions []csv.Question) {
	fmt.Printf("You will be asked %d questions in this test.\n\nPress [Enter] when you are ready.", len(questions))
	_, err := fmt.Scanln()
	if err != nil {
		log.Fatal(err.Error())
	}
}
