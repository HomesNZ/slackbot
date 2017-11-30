package git

import (
	"fmt"

	"github.com/google/go-github/github"
)

func GetRepositoryInformation() []*github.Repository {
	repos, response, err := client.GetRepositories().ListByOrg(ctx, "homesnz", &github.RepositoryListByOrgOptions{})

	fmt.Println("Repo response: ", response.Status)

	if err != nil {
		fmt.Println("error retrieving repositories", err)
		return nil
	}
	return repos
}
