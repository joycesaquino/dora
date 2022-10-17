package types

import (
	"dora/internal/entity"
	"encoding/json"
	"fmt"
	"time"
)

// ReviewRequestEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review
type ReviewRequestEvent struct {
	Action      string            `json:"action"`
	Review      GithubReview      `json:"review"`
	PullRequest GithubPullRequest `json:"pull_request"`
	Repository  GithubRepository  `json:"repository"`
	Sender      GithubUser        `json:"sender"`
}

// PullRequestEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request
type PullRequestEvent struct {
	Action      string            `json:"action"`
	PullRequest GithubPullRequest `json:"pull_request"`
	Repository  GithubRepository  `json:"repository"`
	Sender      GithubUser        `json:"sender"`
}

// PullRequestReviewCommentEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
type PullRequestReviewCommentEvent struct {
	Action      string            `json:"action"`
	Comment     GithubComment     `json:"comment"`
	PullRequest GithubPullRequest `json:"pull_request"`
	Repository  GithubRepository  `json:"repository"`
	Sender      GithubUser        `json:"sender"`
}

type GithubComment struct {
	URL                 string     `json:"url"`
	PullRequestReviewId int        `json:"pull_request_review_id"`
	Id                  int        `json:"id"`
	NodeId              string     `json:"node_id"`
	DiffHunk            string     `json:"diff_hunk"`
	Path                string     `json:"path"`
	Position            int        `json:"position"`
	OriginalPosition    int        `json:"original_position"`
	CommitId            string     `json:"commit_id"`
	OriginalCommitId    string     `json:"original_commit_id"`
	User                GithubUser `json:"user"`
	Body                string     `json:"body"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	HTMLURL             string     `json:"html_url"`
	PullRequestURL      string     `json:"pull_request_url"`
	AuthorAssociation   string     `json:"author_association"`
}

type GithubPullRequest struct {
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
	User               GithubUser    `json:"user"`
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

type GithubReview struct {
	Id                int         `json:"id"`
	NodeId            string      `json:"node_id"`
	User              GithubUser  `json:"user"`
	Body              interface{} `json:"body"`
	CommitId          string      `json:"commit_id"`
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

type GithubRepository struct {
	Id                  int         `json:"id"`
	NodeId              string      `json:"node_id"`
	Name                string      `json:"name"`
	FullName            string      `json:"full_name"`
	Private             bool        `json:"private"`
	Owner               GithubUser  `json:"owner"`
	HTMLURL             string      `json:"html_url"`
	Description         interface{} `json:"description"`
	Fork                bool        `json:"fork"`
	URL                 string      `json:"url"`
	ForksURL            string      `json:"forks_url"`
	KeysURL             string      `json:"keys_url"`
	CollaboratorsURL    string      `json:"collaborators_url"`
	TeamsURL            string      `json:"teams_url"`
	HooksURL            string      `json:"hooks_url"`
	IssueEventsURL      string      `json:"issue_events_url"`
	EventsURL           string      `json:"events_url"`
	AssigneesURL        string      `json:"assignees_url"`
	BranchesURL         string      `json:"branches_url"`
	TagsURL             string      `json:"tags_url"`
	BlobsURL            string      `json:"blobs_url"`
	GitTagsURL          string      `json:"git_tags_url"`
	GitRefsURL          string      `json:"git_refs_url"`
	TreesURL            string      `json:"trees_url"`
	StatusesURL         string      `json:"statuses_url"`
	LanguagesURL        string      `json:"languages_url"`
	StargazersURL       string      `json:"stargazers_url"`
	ContributorsURL     string      `json:"contributors_url"`
	SubscribersURL      string      `json:"subscribers_url"`
	SubscriptionURL     string      `json:"subscription_url"`
	CommitsURL          string      `json:"commits_url"`
	GitCommitsURL       string      `json:"git_commits_url"`
	CommentsURL         string      `json:"comments_url"`
	IssueCommentURL     string      `json:"issue_comment_url"`
	ContentsURL         string      `json:"contents_url"`
	CompareURL          string      `json:"compare_url"`
	MergesURL           string      `json:"merges_url"`
	ArchiveURL          string      `json:"archive_url"`
	DownloadsURL        string      `json:"downloads_url"`
	IssuesURL           string      `json:"issues_url"`
	PullsURL            string      `json:"pulls_url"`
	MilestonesURL       string      `json:"milestones_url"`
	NotificationsURL    string      `json:"notifications_url"`
	LabelsURL           string      `json:"labels_url"`
	ReleasesURL         string      `json:"releases_url"`
	DeploymentsURL      string      `json:"deployments_url"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
	PushedAt            time.Time   `json:"pushed_at"`
	GitURL              string      `json:"git_url"`
	SSHURL              string      `json:"ssh_url"`
	CloneURL            string      `json:"clone_url"`
	SvnURL              string      `json:"svn_url"`
	Homepage            interface{} `json:"homepage"`
	Size                int         `json:"size"`
	StargazersCount     int         `json:"stargazers_count"`
	WatchersCount       int         `json:"watchers_count"`
	Language            string      `json:"language"`
	HasIssues           bool        `json:"has_issues"`
	HasProjects         bool        `json:"has_projects"`
	HasDownloads        bool        `json:"has_downloads"`
	HasWiki             bool        `json:"has_wiki"`
	HasPages            bool        `json:"has_pages"`
	ForksCount          int         `json:"forks_count"`
	MirrorURL           interface{} `json:"mirror_url"`
	Archived            bool        `json:"archived"`
	Disabled            bool        `json:"disabled"`
	OpenIssuesCount     int         `json:"open_issues_count"`
	License             interface{} `json:"license"`
	Forks               int         `json:"forks"`
	OpenIssues          int         `json:"open_issues"`
	Watchers            int         `json:"watchers"`
	DefaultBranch       string      `json:"default_branch"`
	AllowSquashMerge    bool        `json:"allow_squash_merge"`
	AllowMergeCommit    bool        `json:"allow_merge_commit"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge"`
}

type GithubUser struct {
	Login             string `json:"login"`
	Id                int    `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

func (request PullRequestEvent) Converter() *entity.PR {
	return &entity.PR{
		Action: request.Action,
		User: entity.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: entity.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: entity.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: entity.PullRequest{
			Id:                 request.PullRequest.Id,
			URL:                request.PullRequest.URL,
			Title:              request.PullRequest.Title,
			Assignee:           request.PullRequest.Assignee,
			RequestedReviewers: request.PullRequest.RequestedReviewers,
			RequestedTeams:     request.PullRequest.RequestedTeams,
			AuthorAssociation:  request.PullRequest.AuthorAssociation,
			CreatedAt:          request.PullRequest.CreatedAt,
			UpdatedAt:          request.PullRequest.UpdatedAt,
			ClosedAt:           request.PullRequest.ClosedAt,
			MergedAt:           request.PullRequest.MergedAt,
		},
		Review: nil,
	}
}

func (request ReviewRequestEvent) Converter() entity.PR {
	return entity.PR{
		Action: request.Action,
		User: entity.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: entity.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: entity.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: entity.PullRequest{
			Id:                 request.PullRequest.Id,
			URL:                request.PullRequest.URL,
			Title:              request.PullRequest.Title,
			Assignee:           request.PullRequest.Assignee,
			RequestedReviewers: request.PullRequest.RequestedReviewers,
			RequestedTeams:     request.PullRequest.RequestedTeams,
			AuthorAssociation:  request.PullRequest.AuthorAssociation,
			CreatedAt:          request.PullRequest.CreatedAt,
			UpdatedAt:          request.PullRequest.UpdatedAt,
			ClosedAt:           request.PullRequest.ClosedAt,
			MergedAt:           request.PullRequest.MergedAt,
		},
		Review: &entity.Review{
			Id:     request.Review.Id,
			NodeId: request.Review.NodeId,
			User: entity.User{
				Login: request.Review.User.Login,
				Id:    request.Review.User.Id,
			},
			Body:        request.Review.Body,
			CommitId:    request.Review.CommitId,
			SubmittedAt: request.Review.SubmittedAt,
			State:       request.Review.State,
			Comments:    nil,
		},
	}
}

func (request PullRequestReviewCommentEvent) Converter() entity.PR {
	return entity.PR{
		Action: request.Action,
		User: entity.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: entity.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: entity.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: entity.PullRequest{
			Id:                 request.PullRequest.Id,
			URL:                request.PullRequest.URL,
			Title:              request.PullRequest.Title,
			Assignee:           request.PullRequest.Assignee,
			RequestedReviewers: request.PullRequest.RequestedReviewers,
			RequestedTeams:     request.PullRequest.RequestedTeams,
			AuthorAssociation:  request.PullRequest.AuthorAssociation,
			CreatedAt:          request.PullRequest.CreatedAt,
			UpdatedAt:          request.PullRequest.UpdatedAt,
			ClosedAt:           request.PullRequest.ClosedAt,
			MergedAt:           request.PullRequest.MergedAt,
		},
		Review: &entity.Review{
			Comments: []entity.Comment{
				{
					URL:                 request.Comment.URL,
					PullRequestReviewId: request.Comment.PullRequestReviewId,
					Id:                  request.Comment.Id,
					NodeId:              request.Comment.NodeId,
					DiffHunk:            request.Comment.DiffHunk,
					Path:                request.Comment.Path,
					CommitId:            request.Comment.CommitId,
					OriginalCommitId:    request.Comment.CommitId,
					User: entity.User{
						Login: request.Comment.User.Login,
						Id:    request.Comment.User.Id,
					},
					Body:              request.Comment.Body,
					CreatedAt:         request.Comment.CreatedAt,
					UpdatedAt:         request.Comment.UpdatedAt,
					HTMLURL:           request.Comment.HTMLURL,
					AuthorAssociation: request.Comment.AuthorAssociation,
				},
			},
		},
	}
}

func (request PullRequestEvent) toString() {
	b, err := json.MarshalIndent(request, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
