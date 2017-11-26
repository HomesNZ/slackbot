package git

import (
	"context"
	"fmt"
	"os"

	git "../../models"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

//InitGit connects to github
func InitGit() {
	token := os.Getenv("GIT_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client := github.NewClient(tc)

	repo, _, err := client.Repositories.Get(ctx, "homesnz", "homes-web")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &git.Package{
		FullName:    *repo.FullName,
		Description: *repo.Description,
		ForksCount:  *repo.ForksCount,
		StarsCount:  *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)

	fmt.Print("Init git")
}
