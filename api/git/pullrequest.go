package git

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"

	"github.com/google/go-github/github"
	git "github.com/jimsrush/slackbot/models/git"
)

const pullTemplate = `
{{.Title}}
{{.Created}}
{{.URL}}
{{.Days}} days old
`

var opt = &github.PullRequestListOptions{
	State:       "open",
	ListOptions: github.ListOptions{},
}

func GetPullRequestData() string {

	repos := GetRepositoryInformation()

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
	formattedRequests := FormatRequests(rawRequests)

	return strings.Join(formattedRequests, "\n")
}

func FormatRequests(requests []*github.PullRequest) []string {
	var formattedRequests []string

	for _, req := range requests {

		s := git.SimplePR{
			Title:   req.GetTitle(),
			Created: req.CreatedAt.Format("Jan 2, 2006"),
			URL:     req.GetHTMLURL(),
			Days:    calculateDays(req.GetCreatedAt()),
		}

		formatted := FormatPR(s)
		formattedRequests = append(formattedRequests, formatted)
	}
	return formattedRequests
}

func GetPullRequestDataByRepo(repository string) string {
	requests, response, err := client.PullRequests.List(ctx, "homesnz", repository, opt)

	if err != nil {
		return fmt.Sprintf("Something went wrong: %s", response.Status)
	}
	formattedRequests := FormatRequests(requests)
	return strings.Join(formattedRequests, "\n")
}

func FormatPR(s git.SimplePR) string {
	var doc bytes.Buffer
	t := template.New("pullTemplate")
	t, _ = t.Parse(pullTemplate)
	t.Execute(&doc, s)
	st := doc.String()
	return st
}
