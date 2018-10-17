package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const apiHost = "api.stackexchange.com"
const apiVersion = "2.2"
const apiSearchEndpoint = "search"
const defaultSite = "stackoverflow"

type Parameters struct {
	Page      string
	PageSize  string
	FromDate  time.Time
	ToDate    time.Time
	Order     string
	Sort      string
	Tagged    string
	NotTagged string
	InTitle   string
	Site      string
}

type Response struct {
	Items          []Item `json:"items"`
	HasMore        bool   `json:"has_more"`
	QuotaMax       int    `json:"quota_max"`
	QuotaRemaining int    `json:"quota_remaining"`
}

type Item struct {
	Tags             []string `json:"tags"`
	Owner            Owner    `json:"owner"`
	IsAnswered       bool     `json:"is_answered"`
	ViewCount        int      `json:"view_count"`
	FavoriteCount    int      `json:"favorite_count"`
	DownVoteCount    int      `json:"down_vote_count"`
	UpVoteCount      int      `json:"up_vote_count"`
	AnswerCount      int      `json:"answer_count"`
	Score            int      `json:"score"`
	LastActivityDate int      `json:"last_activity_date"`
	CreationDate     int      `json:"creation_date"`
	LastEditDate     int      `json:"last_edit_date"`
	QuestionID       int      `json:"question_id"`
	Link             string   `json:"link"`
	Title            string   `json:"title"`
	Body             string   `json:"body"`
}

type Owner struct {
	Reputation   int    `json:"reputation"`
	UserID       int    `json:"user_id"`
	UserType     string `json:"user_type"`
	AcceptRate   int    `json:"accept_rate"`
	ProfileImage string `json:"profile_image"`
	DisplayName  string `json:"display_name"`
	Link         string `json:"link"`
}

func Search(ctx context.Context, p Parameters, client *http.Client) (Response, error) {
	data := p.urlValues()
	r, err := http.NewRequest("GET", urlBuilder(apiSearchEndpoint), strings.NewReader(data.Encode()))
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

func (p *Parameters) urlValues() url.Values {
	var values = make(url.Values)
	p.set(values, "page", p.Page)
	p.set(values, "pagesize", p.PageSize)
	p.set(values, "order", p.Order)
	p.set(values, "sort", p.Sort)
	p.set(values, "tagged", p.Tagged)
	p.set(values, "nottagged", p.NotTagged)
	p.set(values, "intitle", p.InTitle)
	if p.Site == "" {
		values.Add("site", defaultSite)
	} else {
		values.Add("site", p.Site)
	}
	return values
}

func (p *Parameters) set(values url.Values, key string, value string) {
	if value != "" {
		values.Add(key, value)
	}
}

func urlBuilder(endpoint string) string {
	return "http://" + apiHost + "/" + apiVersion + "/" + endpoint
}
