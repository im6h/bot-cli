package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// httpDevtoCmd represents the httpDevto command
var httpDevtoCmd = &cobra.Command{
	Use:   "devto",
	Short: "A brief description of your command",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		httpDevExecute(cmd, args)
	},
}

type devtoArtile struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Publish     string `json:"published_timestamp"`
	Url         string `json:"url"`
}

func init() {
	rootCmd.AddCommand(httpDevtoCmd)
}

func fetchDevtoArticle() {
	var url string = fmt.Sprintf("https://dev.to/api/articles?page=1&per_page=10&top=2")
	var articles []*devtoArtile

	body := responseData(url)

	err := json.Unmarshal(body, &articles)
	if err != nil {
		log.Panicf("httpDevto - fetchDevtoArticle - error when unmarshaling data: %s\n", err)
	}

	for _, art := range articles {
		article, err := json.MarshalIndent(*art, "", " ")
		if err != nil {
			log.Panicf("httpDevto - fetchDevtoArticle - error when marshaling data: %s\n", err)
		}
		fmt.Printf("%v\n", string(article))
	}
}

func httpDevExecute(cmd *cobra.Command, args []string) {
	fetchDevtoArticle()
}
