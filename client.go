package tranco

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Client is an API client for Tranco API.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Username   string
	Password   string
}

func New(username, password string) *Client {
	return &Client{
		BaseURL:    "https://tranco-list.eu",
		HTTPClient: http.DefaultClient,
		Username:   username,
		Password:   password,
	}
}

//func NewAuthClient(username, password string) (*Client, error) {
//	return &Client{
//		BaseURL:    "https://tranco-list.eu",
//		HTTPClient: http.DefaultClient,
//		Username:   username,
//		Password:   password,
//	}, nil
//}

func (cli *Client) GetRanks(ctx context.Context, domain string) (Ranks, error) {
	if domain == "" {
		return Ranks{}, errors.New("domain must not be empty")
	}
	var api ApiRanks
	err := cli.get(ctx, "/api/ranks/domain/"+domain, &api)
	if err != nil {
		return Ranks{}, fmt.Errorf("fail to get ranks: %w", err)
	}

	var r Ranks
	r.Domain = api.Domain

	if len(api.Ranks) == 0 {
		r.Ranks = []RankEntry{}
		return r, nil
	}

	for _, rank := range api.Ranks {
		r.Ranks = append(r.Ranks, RankEntry{Rank: rank.Rank, Date: rank.Date.Time})
	}

	return r, nil
}

func (cli *Client) GetListMetadataById(ctx context.Context, id string) (ListMetadata, error) {
	if id == "" {
		return ListMetadata{}, errors.New("id must not be empty")
	}
	var api ApiListMetadata
	err := cli.get(ctx, "/api/lists/id/"+id, &api)
	if err != nil {
		return ListMetadata{}, fmt.Errorf("fail to get list by Id: %w", err)
	}

	l := ListMetadata{
		ListId:    api.ListId,
		Available: api.Available,
		Failed:    api.Failed,
		Download:  api.Download,
		CreatedOn: api.CreatedOn.Time,
		Configuration: Configuration{
			Providers:         api.Configuration.Providers,
			ListPrefix:        api.Configuration.ListPrefix,
			EndDate:           api.Configuration.EndDate.Time,
			FilterTLD:         api.Configuration.FilterTLD,
			FilterPLD:         api.Configuration.FilterPLD,
			CombinationMethod: api.Configuration.CombinationMethod,
			StartDate:         api.Configuration.StartDate.Time,
			IsDailyList:       api.Configuration.IsDailyList,
		},
	}

	return l, nil
}

func (cli *Client) GetListMetadataByDate(ctx context.Context, date time.Time) (ListMetadata, error) {
	s := date.Format("20060102")

	var api ApiListMetadata
	err := cli.get(ctx, "/api/lists/date/"+s, &api)
	if err != nil {
		return ListMetadata{}, fmt.Errorf("fail to get list by date: %w", err)
	}

	l := ListMetadata{
		ListId:    api.ListId,
		Available: api.Available,
		Failed:    api.Failed,
		Download:  api.Download,
		CreatedOn: api.CreatedOn.Time,
		Configuration: Configuration{
			Providers:         api.Configuration.Providers,
			ListPrefix:        api.Configuration.ListPrefix,
			EndDate:           api.Configuration.EndDate.Time,
			FilterTLD:         api.Configuration.FilterTLD,
			FilterPLD:         api.Configuration.FilterPLD,
			CombinationMethod: api.Configuration.CombinationMethod,
			StartDate:         api.Configuration.StartDate.Time,
			IsDailyList:       api.Configuration.IsDailyList,
		},
	}

	return l, nil
}

func (cli *Client) GetAuthTest(ctx context.Context) error {
	if cli.Username == "" || cli.Password == "" {
		return errors.New("auth test requires username and password")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cli.BaseURL+"/api/auth/test", nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP request: %w", err)
	}

	req.SetBasicAuth(cli.Username, cli.Password)

	resp, err := cli.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("fail to get auth test: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	return nil
}

func (cli *Client) CreateList(ctx context.Context, configuration Configuration) (CreatedList, error) {
	if cli.Username == "" || cli.Password == "" {
		return CreatedList{}, errors.New("auth test requires username and password")
	}

	jsonData, err := json.Marshal(configuration)
	if err != nil {
		return CreatedList{}, fmt.Errorf("failed to marshall configuration: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, cli.BaseURL+"/api/lists/create", bytes.NewBuffer(jsonData))
	if err != nil {
		return CreatedList{}, fmt.Errorf("cannot create HTTP request: %w", err)
	}

	req.SetBasicAuth(cli.Username, cli.Password)

	resp, err := cli.HTTPClient.Do(req)
	if err != nil {
		return CreatedList{}, fmt.Errorf("fail to create lists: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return CreatedList{}, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	var l CreatedList
	if err := json.NewDecoder(resp.Body).Decode(&l); err != nil {
		return CreatedList{}, fmt.Errorf("cannot parse HTTP body: %w", err)
	}

	return l, nil
}

func (cli *Client) get(ctx context.Context, path string, v interface{}) error {
	reqURL := cli.BaseURL + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("cannot create HTTP request: %w", err)
	}

	resp, err := cli.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return fmt.Errorf("cannot parse HTTP body: %w", err)
	}

	return nil
}
