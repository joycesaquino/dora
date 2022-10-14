package github

import "time"

type PullRequest struct {
	URL                string        `json:"url"`
	Id                 int           `json:"id"`
	NodeId             string        `json:"node_id"`
	HTMLURL            string        `json:"html_url"`
	DiffURL            string        `json:"diff_url"`
	PatchURL           string        `json:"patch_url"`
	IssueURL           string        `json:"issue_url"`
	Number             int           `json:"number"`
	State              string        `json:"state"`
	Locked             bool          `json:"locked"`
	Title              string        `json:"title"`
	User               User          `json:"user"`
	Body               string        `json:"body"`
	CreatedAt          time.Time     `json:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at"`
	ClosedAt           interface{}   `json:"closed_at"`
	MergedAt           interface{}   `json:"merged_at"`
	MergeCommitSha     string        `json:"merge_commit_sha"`
	Assignee           interface{}   `json:"assignee"`
	Assignees          []interface{} `json:"assignees"`
	RequestedReviewers []interface{} `json:"requested_reviewers"`
	RequestedTeams     []interface{} `json:"requested_teams"`
	Labels             []interface{} `json:"labels"`
	Milestone          interface{}   `json:"milestone"`
	CommitsURL         string        `json:"commits_url"`
	ReviewCommentsURL  string        `json:"review_comments_url"`
	ReviewCommentURL   string        `json:"review_comment_url"`
	CommentsURL        string        `json:"comments_url"`
	StatusesURL        string        `json:"statuses_url"`
	AuthorAssociation  string        `json:"author_association"`
}
