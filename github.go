package main

import (
	"context"
	"fmt"
	"time"

	"code.gitea.io/sdk/gitea"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func migrateGithubToGitea(githubOrg, githubToken, giteaHost, giteaToken string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken}, // TODO: GET TOKEN FROM GITHUB
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	giteaclient := gitea.NewClient("https://"+giteaHost+"/", giteaToken) // TODO: CONFIGURE WITH URL of your GITEA INSTANCE AND A TOKEN FROM YOUR GITEA INSTANCE

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	// get all pages of results
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, githubOrg, opt) // TODO: SET WITH NAME OF GITHUB ORG YOU WANT CLONED
		if err != nil {
			fmt.Println(err)
			return
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for i := 0; i < len(allRepos); i++ {
		description := ""
		if allRepos[i].Description != nil { // will throw a nil pointer error if description is passed directly to the below struct
			description = *allRepos[i].Description
		}
		giteaclient.MigrateRepo(gitea.MigrateRepoOption{
			CloneAddr: *allRepos[i].CloneURL,
			UID:       4, // TODO: SET WITH THE ID OF YOUR USER IN GITEA (IN MY CASE 4 is the user id of an org on my gitea instance)
			RepoName:  *allRepos[i].Name,
			// Mirror:      true, // TODO: uncomment this if you want gitea to periodically check for changes
			// Private:     true, // TODO: uncomment this if you want the repo to be private on gitea
			Description: description,
		})
		time.Sleep(100 * time.Millisecond) // THIS IS HERE SO THE GITEA SERVER DOESNT GET HAMMERED WITH REQUESTS
	}
}
