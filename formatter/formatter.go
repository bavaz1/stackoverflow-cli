package formatter

import (
	"fmt"

	"github.com/bavaz1/stackoverflow-cli/search"
	"github.com/jaytaylor/html2text"
)

func List(items []search.Item) string {
	var result string
	for index, item := range items {
		result += fmt.Sprintf("%d. %s\n\t", index+1, item.Title)
		result += fmt.Sprintf("(Views: %d - Owner reputation: %d - Answers: %d)\n\n", item.ViewCount, item.Owner.Reputation, item.AnswerCount)
	}
	return result
}

func OneQuestion(item search.Item) string {
	var result string
	result += fmt.Sprintf("%s\n\t", item.Title)
	result += fmt.Sprintf("(Views: %d - Owner reputation: %d - Answers: %d)\n\n", item.ViewCount, item.Owner.Reputation, item.AnswerCount)
	return result
}

func GetQuestionID(item search.Item) int {
	return item.QuestionID
}

func QuestionBody(items []search.Item) string {
	var result string
	for _, item := range items {
		result += fmt.Sprintf("%s\n", item.Body)
	}
	result, err := html2text.FromString(result, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	return result
}

func AnswersBodys(items []search.Item) string {
	var result string
	for index, item := range items {
		result += fmt.Sprintf("%d. ANSWER\n%s\n\n", index+1, item.Body)
	}
	result, err := html2text.FromString(result, html2text.Options{PrettyTables: true})
	if err != nil {
		panic(err)
	}
	return result
}
