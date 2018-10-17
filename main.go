package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bavaz1/stackoverflow-cli/formatter"
	"github.com/bavaz1/stackoverflow-cli/search"
)

func main() {
	ctx := context.Background()
	params := search.Parameters{
		InTitle: "react",
	}
	client := http.Client{
		Timeout: 4 * time.Second,
	}
	resp, err := search.Search(ctx, params, &client)
	if err != nil {
		panic(err)
	}

	formatted := formatter.List(resp.Items)

	fmt.Print(formatted)
	fmt.Print("Choose a number of the questions: ")
	var chosenQuestionLocalID int
	_, err = fmt.Scanf("%d", &chosenQuestionLocalID)
	if err != nil {
		panic(err)
	}
	fmt.Println()

	chosenQuestion := resp.Items[chosenQuestionLocalID-1]
	chosenQuestionID := formatter.GetQuestionID(chosenQuestion)

	params = search.Parameters{
		Filter: "!9Z(-wwYGT",
	}
	resp, err = search.GetQuestionByID(ctx, chosenQuestionID, params, &client)
	if err != nil {
		panic(err)
	}

	fmt.Print("\033[H\033[2J")
	fmt.Println("***** THE QUESTION *****\n")

	formatted = formatter.QuestionBody(resp.Items)
	fmt.Println(formatted)

	params = search.Parameters{
		Filter: "!9Z(-wzu0T",
	}
	resp, err = search.GetAnswersByID(ctx, chosenQuestionID, params, &client)
	if err != nil {
		panic(err)
	}

	fmt.Println("\n\n***** THE ANSWERS *****\n")

	formatted = formatter.AnswersBodys(resp.Items)
	fmt.Println(formatted)
}
