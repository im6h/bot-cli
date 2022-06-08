package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type quote struct {
	Anime     string `json:"anime" example:"anime"`
	Character string `json:"character" example:"character"`
	Quote     string `json:"quote" example:"quote"`
}

var quoteAnimeCmd = &cobra.Command{
	Use:   "anime",
	Short: "Fetch quote anime",
	Long:  `Use to fetch anime's quote`,
	Run: func(cmd *cobra.Command, args []string) {
		quoteAnimeCmdExecute()
	},
}

func init() {
	rootCmd.AddCommand(quoteAnimeCmd)
}

func fetchRandomQuote() {
	var baseURL string = "https://animechan.vercel.app/api"
	var quote *quote

	resp, err := http.Get(baseURL + "/random")
	if err != nil {
		log.Panicf("error when fetch random quote: %v\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error when read data random quote: %v\n", err)
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Panicf("error when unmarshaling data random quote: %v\n", err)
	}

	fmt.Printf("%s\n \t %s in %s\n", quote.Quote, quote.Character, quote.Anime)
}

func quoteAnimeCmdExecute() {
	fetchRandomQuote()
}
