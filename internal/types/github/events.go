package github

import (
	"dora/internal/types/dora"
)

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

// PullRequestReviewCommentEvent https://docs.github.com/pt/developers/webhooks-and-events/webhooks/webhook-events-and-payloads#pull_request_review_comment
type PullRequestReviewCommentEvent struct {
	Action      string      `json:"action"`
	Comment     Comment     `json:"comment"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
	Sender      User        `json:"sender"`
}

func (request PullRequestEvent) Converter() *dora.PR {
	return &dora.PR{
		Action: request.Action,
		User: dora.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: dora.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: dora.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: dora.PullRequest{
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

func (request ReviewRequestEvent) Converter() dora.PR {
	return dora.PR{
		Action: request.Action,
		User: dora.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: dora.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: dora.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: dora.PullRequest{
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
		Review: &dora.Review{
			Id:     request.Review.Id,
			NodeId: request.Review.NodeId,
			User: dora.User{
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

func (request PullRequestReviewCommentEvent) Converter() dora.PR {
	return dora.PR{
		Action: request.Action,
		User: dora.User{
			Login: request.Sender.Login,
			Id:    request.Sender.Id,
		},
		Repository: dora.Repository{
			Id:       request.Repository.Id,
			Name:     request.Repository.Name,
			FullName: request.Repository.FullName,
			Owner: dora.User{
				Login: request.Repository.Owner.Login,
				Id:    request.Repository.Owner.Id,
			},
		},
		PullRequest: dora.PullRequest{
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
		Review: &dora.Review{
			Comments: []Comment{
				request.Comment,
			},
		},
	}
}
