package custommock

import (
	"context"

	"github.com/zawa-t/pr-reviewer/commentator/src/platform/github"
	mock "github.com/zawa-t/pr-reviewer/commentator/src/test/mock/github"
)

var DefaultGithubReview = &mock.ClientMock{
	CreateCommentFunc: func(ctx context.Context, data github.CommentData) error {
		return nil
	},
}
