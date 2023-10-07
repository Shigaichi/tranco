package tranco

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetRanks(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name           string
		args           args
		want           Ranks
		testDataPath   string
		testStatusCode int
		wantErr        bool
	}{
		{
			name: "ok",
			args: args{domain: "example.com"},
			want: Ranks{
				Ranks: []RankEntry{
					{Date: parseDate("2023-09-02"), Rank: 192},
					{Date: parseDate("2023-09-01"), Rank: 192},
					{Date: parseDate("2023-08-31"), Rank: 191},
					{Date: parseDate("2023-08-30"), Rank: 191},
					{Date: parseDate("2023-08-29"), Rank: 189},
					{Date: parseDate("2023-08-28"), Rank: 188},
					{Date: parseDate("2023-08-27"), Rank: 190},
					{Date: parseDate("2023-08-26"), Rank: 191},
					{Date: parseDate("2023-08-25"), Rank: 190},
					{Date: parseDate("2023-08-24"), Rank: 190},
					{Date: parseDate("2023-08-23"), Rank: 190},
					{Date: parseDate("2023-08-22"), Rank: 189},
					{Date: parseDate("2023-08-21"), Rank: 188},
					{Date: parseDate("2023-08-20"), Rank: 190},
					{Date: parseDate("2023-08-19"), Rank: 191},
					{Date: parseDate("2023-08-18"), Rank: 191},
					{Date: parseDate("2023-08-17"), Rank: 191},
					{Date: parseDate("2023-08-16"), Rank: 191},
					{Date: parseDate("2023-08-15"), Rank: 191},
					{Date: parseDate("2023-08-14"), Rank: 191},
					{Date: parseDate("2023-08-13"), Rank: 192},
					{Date: parseDate("2023-08-12"), Rank: 192},
					{Date: parseDate("2023-08-11"), Rank: 192},
					{Date: parseDate("2023-08-10"), Rank: 191},
					{Date: parseDate("2023-08-09"), Rank: 191},
					{Date: parseDate("2023-08-08"), Rank: 192},
					{Date: parseDate("2023-08-07"), Rank: 191},
					{Date: parseDate("2023-08-06"), Rank: 191},
					{Date: parseDate("2023-08-05"), Rank: 192},
					{Date: parseDate("2023-08-04"), Rank: 196},
					{Date: parseDate("2023-08-03"), Rank: 201},
					{Date: parseDate("2023-08-02"), Rank: 204},
					{Date: parseDate("2023-08-01"), Rank: 206},
					{Date: parseDate("2023-07-31"), Rank: 836},
					{Date: parseDate("2023-07-30"), Rank: 838},
					{Date: parseDate("2023-07-29"), Rank: 839},
					{Date: parseDate("2023-07-28"), Rank: 843},
					{Date: parseDate("2023-07-27"), Rank: 848},
					{Date: parseDate("2023-07-26"), Rank: 851},
					{Date: parseDate("2023-07-25"), Rank: 855},
					{Date: parseDate("2023-07-24"), Rank: 853},
				},
				Domain: "example.com",
			},
			testDataPath:   "testdata/get_ranks1.json",
			testStatusCode: 200,
			wantErr:        false,
		},
		{
			name: "not exist domain returns empty result",
			args: args{domain: "no-example.com"},
			want: Ranks{
				Ranks:  []RankEntry{},
				Domain: "no-example.com",
			},
			testDataPath:   "testdata/get_ranks2.json",
			testStatusCode: 200,
			wantErr:        false,
		},
		{
			name:           "returns error if domain is empty",
			args:           args{domain: ""},
			want:           Ranks{},
			testDataPath:   "",
			testStatusCode: 200,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, clean := setup(t, tt.testDataPath, tt.testStatusCode, http.MethodGet, fmt.Sprintf("/api/ranks/domain/%s", tt.args.domain))
			defer clean()
			cli := &Client{
				httpClient: http.DefaultClient,
				baseUrl:    server.baseUrl,
			}
			got, err := cli.GetRanks(context.Background(), tt.args.domain)
			if (err != nil) != tt.wantErr {
				t.Errorf("api error %v", err)
				return
			} else if tt.wantErr {
				if err == nil {
					t.Error("error was expected but not occurred")
				}
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("response is mimatch:\n%s", diff)
			}
		})
	}
}

func TestClient_GetListMetadataById(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name           string
		args           args
		want           ListMetadata
		testDataPath   string
		testStatusCode int
		wantErr        bool
	}{
		{
			name: "ok",
			args: args{id: "JX5LY"},
			want: ListMetadata{
				ListId:    "JX5LY",
				Available: true,
				Failed:    false,
				Download:  "https://tranco-list.eu/download/JX5LY/1000000",
				CreatedOn: time.Date(2022, 12, 11, 22, 0, 9, 199647000, time.UTC),
				Configuration: Configuration{
					Providers:         []string{"alexa", "umbrella", "majestic", "farsight"},
					ListPrefix:        "full",
					EndDate:           parseDate("2022-12-11"),
					FilterTLD:         "false",
					FilterPLD:         "on",
					CombinationMethod: "dowdall",
					StartDate:         parseDate("2022-11-12"),
					IsDailyList:       true,
				},
			},
			testDataPath:   "testdata/get_list_metadata1.json",
			testStatusCode: 200,
			wantErr:        false,
		},
		{
			name:           "returns error if id is not exist",
			args:           args{id: "xxx"},
			want:           ListMetadata{},
			testDataPath:   "testdata/get_list_metadata2.json",
			testStatusCode: 404,
			wantErr:        true,
		},
		{
			name:           "returns error if id is empty",
			args:           args{id: ""},
			want:           ListMetadata{},
			testDataPath:   "",
			testStatusCode: 200,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, clean := setup(t, tt.testDataPath, tt.testStatusCode, http.MethodGet, fmt.Sprintf("/api/lists/id/%s", tt.args.id))
			defer clean()
			cli := &Client{
				httpClient: http.DefaultClient,
				baseUrl:    server.baseUrl,
			}
			got, err := cli.GetListMetadataById(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListMetadataById() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				if err == nil {
					t.Error("error was expected but not occurred")
				}
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("response is mimatch:\n%s", diff)
			}
		})
	}
}

func TestClient_GetListMetadataByDate(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name                string
		args                args
		want                ListMetadata
		testDataPath        string
		testStatusCode      int
		expectedRequestPath string
		wantErr             bool
	}{
		{
			name: "ok",
			args: args{date: parseDate("2022-12-11")},
			want: ListMetadata{
				ListId:    "JX5LY",
				Available: true,
				Failed:    false,
				Download:  "https://tranco-list.eu/download/JX5LY/1000000",
				CreatedOn: time.Date(2022, 12, 11, 22, 0, 9, 199647000, time.UTC),
				Configuration: Configuration{
					Providers:         []string{"alexa", "umbrella", "majestic", "farsight"},
					ListPrefix:        "full",
					EndDate:           parseDate("2022-12-11"),
					FilterTLD:         "false",
					FilterPLD:         "on",
					CombinationMethod: "dowdall",
					StartDate:         parseDate("2022-11-12"),
					IsDailyList:       true,
				},
			},
			testDataPath:        "testdata/get_list_metadata1.json",
			testStatusCode:      200,
			expectedRequestPath: "20221211",
			wantErr:             false,
		},
		{
			name:                "returns error if id is not exist",
			args:                args{date: parseDate("1192-12-11")},
			want:                ListMetadata{},
			testDataPath:        "testdata/get_list_metadata2.json",
			testStatusCode:      404,
			expectedRequestPath: "11921211",
			wantErr:             true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server, clean := setup(t, tt.testDataPath, tt.testStatusCode, http.MethodGet, fmt.Sprintf("/api/lists/date/%s", tt.expectedRequestPath))
			defer clean()
			cli := &Client{
				httpClient: http.DefaultClient,
				baseUrl:    server.baseUrl,
			}
			got, err := cli.GetListMetadataByDate(context.Background(), tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListMetadataById() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if tt.wantErr {
				if err == nil {
					t.Error("error was expected but not occurred")
				}
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("response is mimatch:\n%s", diff)
			}
		})
	}
}

func setup(t *testing.T, mockResponseBodyFile string, mockStatusCode int, expectedMethod, expectedRequestPath string) (*Client, func()) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != expectedMethod {
			t.Fatalf("request method is mismatch. got = %s, want %s", req.Method, expectedMethod)
		}
		if req.URL.Path != expectedRequestPath {
			t.Fatalf("request path is mismatch. got = %s, want %s", req.URL.Path, expectedRequestPath)
		}

		w.WriteHeader(mockStatusCode)

		if mockResponseBodyFile == "" {
			return
		}

		bodyBytes, err := os.ReadFile(mockResponseBodyFile)
		if err != nil {
			t.Fatalf("failed to read testdata (%s). err %s", mockResponseBodyFile, err.Error())
		}
		w.Write(bodyBytes)
	}))

	cli := &Client{
		baseUrl:    server.URL,
		httpClient: server.Client(),
	}
	teardown := func() {
		server.Close()
	}

	return cli, teardown
}

func parseDate(dateStr string) time.Time {
	date, _ := time.Parse("2006-01-02", dateStr)
	return date
}
