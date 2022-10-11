package github

// ReviewRequestEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
type ReviewRequestEvent struct {
	Action      string      `json:"action"`
	Review      Review      `json:"review"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      User        `json:"sender"`
}

// PullRequestEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
type PullRequestEvent struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      User        `json:"sender"`
}
