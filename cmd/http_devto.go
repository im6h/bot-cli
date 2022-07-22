package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// httpDevtoCmd represents the httpDevto command
var httpDevtoCmd = &cobra.Command{
	Use:   "devto",
	Short: "A brief description of your command",
	Long:  "Fetch dev.to article from api with flag",
	Run: func(cmd *cobra.Command, args []string) {
		httpDevExecute(cmd, args)
	},
}

type user struct {
	Name            string `json:"name"`
	TwitterUsername string `json:"twitter_username"`
	GithubUsername  string `json:"github_username"`
	WebsiteUrl      string `json:"website_url"`
}
type devtoArtile struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Publish     string `json:"published_timestamp"`
	Url         string `json:"url"`
	Tags        string `json:"tags"`
	User        user   `json:"user"`
}

func init() {
	rootCmd.AddCommand(httpDevtoCmd)
}

func fetchDevtoTopArticle(page, per_page string) {
	var url string = fmt.Sprintf("https://dev.to/api/articles?page=%s&per_page=%s", page, per_page)
	var articles []*devtoArtile

	body := responseData(url)

	err := json.Unmarshal(body, &articles)
	if err != nil {
		log.Panicf("httpDevto - fetchDevtoArticle - error when unmarshaling data: %s\n", err)
	}

	mapArticles := lo.Map(articles, func(x *devtoArtile, _ int) struct {
		Id          int    `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Publish     string `json:"published_timestamp"`
		Url         string `json:"url"`
		Tags        string `json:"tags"`
		Username    string `json:"username"`
		Github      string `json:"github"`
	} {
		return struct {
			Id          int    `json:"id"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Publish     string `json:"published_timestamp"`
			Url         string `json:"url"`
			Tags        string `json:"tags"`
			Username    string `json:"username"`
			Github      string `json:"github"`
		}{
			Id:          x.Id,
			Title:       x.Title,
			Description: x.Description,
			Publish:     x.Publish,
			Url:         x.Url,
			Tags:        x.Tags,
			Username:    x.User.Name,
			Github:      fmt.Sprintf("https://github.com/%s", x.User.GithubUsername),
		}
	})

	for _, art := range mapArticles {
		article, err := json.MarshalIndent(art, "", "  ")
		if err != nil {
			log.Panicf("httpDevto - fetchDevtoArticle - error when marshaling data: %s\n", err)
		}
		fmt.Printf("%v\n", string(article))
	}
}

func httpDevExecute(cmd *cobra.Command, args []string) {
	var page string = "1"
	var perPage string = "10"

	if len(args) > 0 {
		if args[0] != "" {
			page = args[0]
		}
		if args[1] != "" {
			perPage = args[1]
		}
	}

	fetchDevtoTopArticle(page, perPage)

}
