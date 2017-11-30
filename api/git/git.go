package git

import (
	"context"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client github.Client
var repos []*github.Repository
var ctx context.Context

const pullTemplate = `
{{.Title}}
{{.Created}}
{{.URL}}
{{.Days}} days old
`

//InitGit connects to github
func InitGit() {
	token := os.Getenv("GIT_TOKEN")

	ctx = context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client = *github.NewClient(tc)
}

func calculateDays(then time.Time) int {
	duration := time.Now().Sub(then)
	return int(duration.Hours() / 24)
}
