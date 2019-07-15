package qns

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

func Table(results []Result) {
	fmt.Println("\nYou can see your detailed results below:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Your answer", "Correct answer", "Result", "Points earned"})
	for i, result := range results {
		if result.Correct {
			table.Append([]string{strconv.Itoa(i+1), result.UserAnswer, result.CorrectAnswer, "Right", strconv.Itoa(result.Points)})
		} else {
			table.Append([]string{strconv.Itoa(i+1), result.UserAnswer, result.CorrectAnswer, "Wrong", strconv.Itoa(result.Points)})
		}
	}

	table.Render()
}
