package git

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client github.Client
var repos []*github.Repository

const pullTemplate = `
Title: ".Title"
Date: ".Created"
URL: ".URL"
`

//InitGit connects to github
func InitGit() {
	token := os.Getenv("GIT_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client = *github.NewClient(tc)

	// o := &github.RepositoryListByOrgOptions{}

	repos, _, _ = client.GetRepositories().ListByOrg(ctx, "homesnz", &github.RepositoryListByOrgOptions{})
	for _, repo := range repos {
		fmt.Println("Repo is:", repo.GetName())
	}

	opt := &github.PullRequestListOptions{
		State:       "open",
		ListOptions: github.ListOptions{},
	}

	requests, _, err := client.PullRequests.List(ctx, "homesnz", "homes-web", opt)

	for _, req := range requests {
		s := SimplePR{
			Title:   req.GetTitle(),
			Created: req.GetCreatedAt(),
			URL:     req.GetURL(),
		}
		// fmt.Println("Simple object is simple", s)
		formatted := FormatPR(s)
		fmt.Println("Simple PR:", formatted)
		fmt.Println("Simple PR:", s)

		// fmt.Println("request is", req.GetTitle())
		// fmt.Println("request is", req.GetCreatedAt())
		// fmt.Println("request is", req.GetURL())
		// fmt.Println("request is", req.GetURL())
	}

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
}

type SimplePR struct {
	Title   string
	Created time.Time
	URL     string
}

func FormatPR(s SimplePR) string {
	var doc bytes.Buffer
	t := template.New("PullRequest")
	t, _ = t.Parse(pullTemplate)
	t.Execute(&doc, s)
	st := doc.String()
	return st
}
