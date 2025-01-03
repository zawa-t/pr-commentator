package dependency

import (
	"log/slog"
	"os"

	bitbucketClient "github.com/zawa-t/pr/src/platform/bitbucket/client"

	githubClient "github.com/zawa-t/pr/src/platform/github/client"

	"github.com/zawa-t/pr/src/report"
	"github.com/zawa-t/pr/src/report/role"

	"github.com/zawa-t/pr/src/platform/http"
)

func NewReporter(roleNum int) (reporter report.Reporter) {
	switch roleNum {
	case role.LocalComment:
		reporter = role.NewLocalCommentator()
	case role.BitbucketPRComment:
		reporter = role.NewBitbucketPRCommentator(bitbucketClient.NewCustomClient(http.NewClient()))
	case role.GithubPRComment:
		reporter = role.NewGithubPRCommentator(githubClient.NewCustomClient(http.NewClient()))
	// case role.GithubPRCheck:
	// 	reporter = role.NewGithubPRChecker(githubClient.NewCustomClient(http.NewClient()))
	case role.GithubCheck:
		reporter = role.NewGithubChecker(githubClient.NewCustomClient(http.NewClient()))
	default:
		slog.Error("Unsupported role was set.")
		os.Exit(1)
	}
	return
}
