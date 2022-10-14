package dora

import (
	"dora/internal/types/github"
	"time"
)

type User struct {
	Login string `json:"login"`
	Id    int    `json:"id"`
}

type Repository struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    User   `json:"owner"`
}

type PullRequest struct {
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

type Review struct {
	Id          int              `json:"id"`
	NodeId      string           `json:"node_id"`
	User        User             `json:"user"`
	Body        interface{}      `json:"body"`
	CommitId    string           `json:"commit_id"`
	SubmittedAt time.Time        `json:"submitted_at"`
	State       string           `json:"state"`
	Comments    []github.Comment `json:"comments"`
}

type PR struct {
	Action      string      `json:"action"`
	User        User        `json:"user"`
	Repository  Repository  `json:"repository"`
	PullRequest PullRequest `json:"pull_request"`
	Review      *Review     `json:"review"`
}
