package cmd

import (
	"log"
	"strconv"

	"context"
	"os"

	"github.com/fatih/color"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var client *github.Client

func init() {
	ctx = context.Background()
	token = os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("githubApiExecute - Error: Unauthorized: No token present")
	}

	rootCmd.AddCommand(githubApiCmd)

	githubApiCmd.AddCommand(githubCmdCreate)
	githubApiCmd.AddCommand(githubCmdDelete)

	githubApiCmd.PersistentFlags().StringVar(&name, "name", "", "Repo name")
	githubApiCmd.PersistentFlags().StringVar(&description, "description", "", "Repo description")
	githubApiCmd.PersistentFlags().BoolVar(&isPrivate, "private", false, "Repo public or private")
	githubApiCmd.PersistentFlags().BoolVar(&isAutoInit, "init", false, "Pass true to create an initial commit with empty README.")

}

var githubApiCmd = &cobra.Command{
	Use:   "ph",
	Short: "A simple gitub cli to repository",
	Run: func(cmd *cobra.Command, args []string) {
		client = authorization(token)
	},
}

var githubCmdCreate = &cobra.Command{
	Use:   "create",
	Short: "create github repository",
	Run: func(cmd *cobra.Command, args []string) {
		client = authorization(token)
		if name == "" {
			log.Fatalf("Name repo is required!")
		}

		repo := &github.Repository{
			Name:        &name,
			Description: &description,
			Private:     &isPrivate,
			AutoInit:    &isAutoInit,
		}

		createRepo(client, repo)
	},
}

var githubCmdDelete = &cobra.Command{
	Use:   "delete",
	Short: "delete github repository",
	Run: func(cmd *cobra.Command, args []string) {
		client = authorization(token)
		if name == "" {
			log.Fatalf("Name repo is required!")
		}

		deleteRepo(client, name)
	},
}

func authorization(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Panicf("Unauthorize token: %s", err)
	}

	rate := strconv.Itoa(resp.Rate.Limit)
	log.Printf("Rate limit: %s\n", color.RedString(rate))

	timeExpiration := resp.TokenExpiration.Time.Format("2006.01.02 15:04:05")

	if !resp.TokenExpiration.IsZero() {
		log.Printf("Time token expiration: %v\n", color.YellowString(timeExpiration))
	}

	log.Printf("Login by username: %s\n", user.GetLogin())
	return client
}

func createRepo(client *github.Client, repo *github.Repository) {
	data, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		log.Fatalf("createRepo - Error: %v", err)
	}

	log.Printf("Repository name: %s\n", color.GreenString(data.GetName()))
	log.Printf("Repository url: %s\n", color.GreenString(data.GetGitURL()))
}

func deleteRepo(client *github.Client, repoName string) {
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Panicf("Unauthorize token: %s", err)
	}

	_, err = client.Repositories.Delete(ctx, user.GetLogin(), repoName)
	if err != nil {
		log.Fatalf("deleteRepo - Error: %v\n", err)
	}

	log.Printf("Deleted repo: %s\n", color.GreenString(repoName))
}
