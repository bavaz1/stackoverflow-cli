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
}
