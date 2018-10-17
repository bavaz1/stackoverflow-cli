package search

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestParametersUrlValues(t *testing.T) {
	p := Parameters{
		Page:    "12",
		InTitle: "react",
	}
	values := p.urlValues()
	expected := "intitle=react&page=12"
	encoded := values.Encode()
	if encoded != expected {
		t.Errorf("Query should be %s, instead of %s", expected, encoded)
	}
}

func TestSearch(t *testing.T) {
	p := Parameters{
		InTitle: "react",
	}
	_, err := Search(context.Background(), p, http.DefaultClient)

	fmt.Println(err)
}
