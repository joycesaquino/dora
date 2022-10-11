package github

import "time"

type Review struct {
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	User              User        `json:"user"`
	Body              interface{} `json:"body"`
	CommitID          string      `json:"commit_id"`
	SubmittedAt       time.Time   `json:"submitted_at"`
	State             string      `json:"state"`
	HTMLURL           string      `json:"html_url"`
	PullRequestURL    string      `json:"pull_request_url"`
	AuthorAssociation string      `json:"author_association"`
	Links             struct {
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		PullRequest struct {
			Href string `json:"href"`
		} `json:"pull_request"`
	} `json:"_links"`
}
