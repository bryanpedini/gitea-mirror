package main

import (
	"fmt"
	"os"

	"github.com/bryanpedini/gitea-mirror-gitea/utils"
)

func main() {
	commandLineArgumentsSlice := os.Args[1:]
	commandLineArguments := utils.StringSliceToMap(commandLineArgumentsSlice)

	if _, ok := commandLineArguments["--github"]; ok {
		githubOrg, githubToken, giteaHost, giteaToken := "", "", "", ""
		migrateGithubToGitea(githubOrg, githubToken, giteaHost, giteaToken)
	} else {
		fmt.Println("Usage: " + os.Args[0] + " [--github]")
	}
}
