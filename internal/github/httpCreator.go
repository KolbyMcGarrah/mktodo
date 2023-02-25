package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const BaseURL = "https://api.github.com/repos/%s/%s/issues"

type HTTPCreator struct {
	client http.Client
}

func NewHTTPCreator() *HTTPCreator {
	return &HTTPCreator{
		client: *http.DefaultClient,
	}
}

func (hc *HTTPCreator) CreateIssue(ctx context.Context, r Request) (string, error) {
	reqBody, _ := json.Marshal(r)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf(BaseURL, r.Owner, r.Repo), bytes.NewBuffer(reqBody))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", strings.TrimSpace(r.GHToken)))
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	res, err := hc.client.Do(req)
	if err != nil {
		return "", err
	}
	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	return resp.URL, err
}
