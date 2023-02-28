package github_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/kolbymcgarrah/mktodo/internal/github"
	"github.com/stretchr/testify/assert"
)

func TestCreateIssue(t *testing.T) {
	type testCase struct {
		name             string
		request          github.Request
		client           *http.Client
		expectedURL      string
		expectedErrorMsg string
	}

	testCases := []testCase{
		{
			name:    "success test",
			request: github.Request{},
			client: &http.Client{Transport: RoundTripFunc(func(*http.Request) *http.Response {
				dataBytes, _ := json.Marshal(github.Response{URL: "testurl"})
				body := io.NopCloser(bytes.NewBuffer(dataBytes))
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       body,
				}
			})},
			expectedURL:      "testurl",
			expectedErrorMsg: "",
		},
		{
			name:    "Unauthorized test",
			request: github.Request{},
			client: &http.Client{Transport: RoundTripFunc(func(*http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusUnauthorized,
				}
			})},
			expectedURL:      "",
			expectedErrorMsg: "unauthorized",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			issueCreator := github.NewHTTPCreator(tc.client)
			url, err := issueCreator.CreateIssue(context.Background(), tc.request)
			assert.Equal(t, tc.expectedURL, url)
			if tc.expectedErrorMsg == "" {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.expectedErrorMsg, err.Error())
			}
		})
	}
}

type RoundTripFunc func(r *http.Request) *http.Response

func (rt RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt(req), nil
}
