package cmd

import (
	"fmt"
	"log"
	"strconv"

	"context"
	"os"

	"github.com/fatih/color"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var ctx context.Context
var name, description string
var isPrivate, isAutoInit bool

var githubApiCmd = &cobra.Command{
	Use:   "ph",
	Short: "A simple gitub cli to repository",
	Run: func(cmd *cobra.Command, args []string) {
		githubApiExecute(cmd, args)
	},
}

func init() {
	ctx = context.Background()
	rootCmd.AddCommand(githubApiCmd)

	githubApiCmd.PersistentFlags().StringVar(&name, "name", "", "Repo name")
	githubApiCmd.PersistentFlags().StringVar(&description, "description", "", "Repo description")
	githubApiCmd.PersistentFlags().BoolVar(&isPrivate, "private", false, "Repo public or private")
	githubApiCmd.PersistentFlags().BoolVar(&isAutoInit, "init", false, "Pass true to create an initial commit with empty README.")

}

func authorization(token string) (*github.Client, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("authorization - GetUser - Error: %v\n", err)
	}

	rate := strconv.Itoa(resp.Rate.Limit)
	log.Printf("Rate limit: %s\n", color.RedString(rate))

	timeExpiration := resp.TokenExpiration.Time.Format("2006.01.02 15:04:05")

	if !resp.TokenExpiration.IsZero() {
		log.Printf("Time token expiration: %v\n", color.YellowString(timeExpiration))
	}

	log.Printf("Login by username: %s\n", *user.Name)
	return client, nil
}

func createNewRepo(client *github.Client, repo *github.Repository) {
	data, _, err := client.Repositories.Create(ctx, "", repo)
	if err != nil {
		log.Fatalf("createNewRepo - Error: %v", err)
	}

	log.Printf("Successfully created new repo: %v\n", color.GreenString(data.GetName()))
	log.Printf("Repository name: %s\n", color.GreenString(data.GetName()))
	log.Printf("Repository url: %s\n", color.GreenString(data.GetCloneURL()))
}

/*
TODO: function delete repo with authorize token
*/

/*
TODO: function fetch trending repo with authorize token
*/

func githubApiExecute(cmd *cobra.Command, args []string) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("githubApiExecute - Error: Unauthorized: No token present")
	}

	client, err := authorization(token)
	if err != nil {
		log.Fatalf("githubApiExecute - Error: %v", err)
	}

	repo := &github.Repository{
		Name:        &name,
		Description: &description,
		Private:     &isPrivate,
		AutoInit:    &isAutoInit,
	}

	createNewRepo(client, repo)
}
