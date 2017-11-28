package git

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"
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
	go GetPullRequestData()
}

func calculateDays(then time.Time) int {
	duration := time.Now().Sub(then)
	return int(duration.Hours() / 24)
}

func GetRepositoryInformation() []*github.Repository {
	repos, response, err := client.GetRepositories().ListByOrg(ctx, "homesnz", &github.RepositoryListByOrgOptions{})

	fmt.Println("Repo response: ", response.Status)

	if err != nil {
		fmt.Println("error retrieving repositories", err)
		return nil
	}
	return repos
}

func GetPullRequestData() string {

	repos := GetRepositoryInformation()

	opt := &github.PullRequestListOptions{
		State:       "open",
		ListOptions: github.ListOptions{},
	}

	var rawRequests []*github.PullRequest

	for _, r := range repos {
		requests, _, err := client.PullRequests.List(ctx, "homesnz", r.GetName(), opt)
		if err != nil {
			fmt.Println("Error returning pull requests: ", err)
		}

		for _, raw := range requests {
			rawRequests = append(rawRequests, raw)
		}
	}

	var formattedRequests []string

	for _, req := range rawRequests {

		s := SimplePR{
			Title:   req.GetTitle(),
			Created: req.CreatedAt.Format("Jan 2, 2006"),
			URL:     req.GetHTMLURL(),
			Days:    calculateDays(req.GetCreatedAt()),
		}

		formatted := FormatPR(s)
		formattedRequests = append(formattedRequests, formatted)
		fmt.Println("Formatted\n", formatted)
	}

	return strings.Join(formattedRequests, "\n")
}

type SimplePR struct {
	Title   string
	Created string
	URL     string
	Days    int
}

func FormatPR(s SimplePR) string {
	var doc bytes.Buffer
	t := template.New("pullTemplate")
	t, _ = t.Parse(pullTemplate)
	t.Execute(&doc, s)
	st := doc.String()
	return st
}
