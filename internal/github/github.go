package github

import "context"

type Request struct {
	GHToken string   `json:"-"`
	Owner   string   `json:"-"`
	Repo    string   `json:"-"`
	Body    string   `json:"body"`
	Title   string   `json:"title"`
	Labels  []string `json:"labels,omitempty"`
}

type Response struct {
	URL string `json:"html_url"`
}

type IssueCreator interface {
	CreateIssue(ctx context.Context, r Request) error
}
