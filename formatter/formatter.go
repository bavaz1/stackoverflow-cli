package formatter

import (
	"fmt"

	"github.com/bavaz1/stackoverflow-cli/search"
)

func List(items []search.Item) string {
	var result string
	for index, item := range items {
		result += fmt.Sprintf("%d. %s\n\t", index+1, item.Title)
		result += fmt.Sprintf("(Views: %d - Owner reputation: %d - Answers: %d)\n\n", item.ViewCount, item.Owner.Reputation, item.AnswerCount)
	}
	return result
}
