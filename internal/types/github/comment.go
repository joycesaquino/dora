package github

import "time"

type Comment struct {
	URL                 string    `json:"url"`
	PullRequestReviewId int       `json:"pull_request_review_id"`
	Id                  int       `json:"id"`
	NodeId              string    `json:"node_id"`
	DiffHunk            string    `json:"diff_hunk"`
	Path                string    `json:"path"`
	Position            int       `json:"position"`
	OriginalPosition    int       `json:"original_position"`
	CommitId            string    `json:"commit_id"`
	OriginalCommitId    string    `json:"original_commit_id"`
	User                User      `json:"user"`
	Body                string    `json:"body"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	HTMLURL             string    `json:"html_url"`
	PullRequestURL      string    `json:"pull_request_url"`
	AuthorAssociation   string    `json:"author_association"`
}
