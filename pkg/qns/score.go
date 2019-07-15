package qns

import "fmt"

func Score(results []Result, correctAnswersNum, totalScore int)  {
	fmt.Printf("\nYou answered %d questions correctly. You have earned %d points.\n", correctAnswersNum, totalScore)

	Table(results)
}
