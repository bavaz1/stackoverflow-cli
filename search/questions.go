package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetQuestionByID(ctx context.Context, id int, p Parameters, client *http.Client) (Response, error) {
	const apiAnswersEndpoint = "questions/%d"
	url := urlBuilder(fmt.Sprintf(apiAnswersEndpoint, id))
	data := p.urlValues()

	r, err := http.NewRequest("GET", url, strings.NewReader(data.Encode()))
	if err != nil {
		return Response{}, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	r = r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	if resp.StatusCode == 200 {
		var result Response
		err = json.Unmarshal(body, &result)
		return result, err
	}

	return Response{}, fmt.Errorf("Undefined status code: %s", string(body))
}
