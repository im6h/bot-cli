package cmd

import (
	"fmt"
	"log"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

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

// httpDevtoCmd represents the httpDevto command
var httpDevtoCmd = &cobra.Command{
	Use:   "devto",
	Short: "Get article in dev.to",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var httpDevtoTopArticleCmd = &cobra.Command{
	Use:   "top",
	Short: "Get top article in dev.to",
	Run: func(cmd *cobra.Command, args []string) {
		fetchDevtoTopArticle(page, limit)
	},
}

func init() {
	rootCmd.AddCommand(httpDevtoCmd)

	// sub-command
	httpDevtoCmd.AddCommand(httpDevtoTopArticleCmd)

	// binding flag
	httpDevtoTopArticleCmd.PersistentFlags().StringVar(&page, "page", "1", "Pagination page")
	httpDevtoTopArticleCmd.PersistentFlags().StringVar(&limit, "limit", "10", "Limit for per page")
}

func fetchDevtoTopArticle(page, per_page string) {
	var url string = fmt.Sprintf("https://dev.to/api/articles?page=%s&per_page=%s", page, per_page)
	var articles []*devtoArtile

	body := ResponseData(url)

	err := json.Unmarshal(body, &articles)
	if err != nil {
		log.Panicf("httpDevto - fetchDevtoArticle - Error: %s\n", err)
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
			log.Panicf("httpDevto - fetchDevtoArticle - Error: %s\n", err)
		}
		fmt.Printf("%v\n", string(article))
	}
}
