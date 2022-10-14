package dora

import "time"

type User struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

type Dora struct {
	Action     string `json:"action"`
	User       User   `json:"user"`
	Repository struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Owner    User   `json:"owner"`
	}
	PullRequest struct {
		Id                 int           `json:"id"`
		URL                string        `json:"url"`
		Title              string        `json:"title"`
		Assignee           interface{}   `json:"assignee"`
		RequestedReviewers []interface{} `json:"requested_reviewers"`
		RequestedTeams     []interface{} `json:"requested_teams"`
		AuthorAssociation  string        `json:"author_association"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		ClosedAt           interface{}   `json:"closed_at"`
		MergedAt           interface{}   `json:"merged_at"`
	}
}
