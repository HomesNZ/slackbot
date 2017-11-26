package git

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client github.Client

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

	lo := github.ListOptions{
		Page:    1,
		PerPage: 1,
	}
	opt := &github.PullRequestListOptions{
		State:       "open",
		ListOptions: lo,
	}

	repo, response, err := client.PullRequests.List(ctx, "homesnz", "homes-web", opt)

	fmt.Println("REsponse is", response)
	fmt.Println("REsponse is", repo)

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}
}

func GetPullRequests() {

}
