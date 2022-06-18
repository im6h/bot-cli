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
		quoteAnimeCmdExecute(args)
	},
}

func init() {
	rootCmd.AddCommand(quoteAnimeCmd)
}

func responseData(url string) []byte {

	resp, err := http.Get(url)
	if err != nil {
		log.Panicf("error when fetch random quote: %v\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("error when read data random quote: %v\n", err)
	}

	return body
}

func fetchQuoteByCharacterName(name string, page string) {
	var url string = fmt.Sprintf("https://animechan.vercel.app/api/quotes/character?name=%s", name)
	var quotes []*quote

	body := responseData(url)

	err := json.Unmarshal(body, &quotes)
	if err != nil {
		log.Panicf("error when unmarshaling data in fetchQuoteByCharacterName: %v\n", err)
	}

	for _, quote := range quotes {
		fmt.Printf(`"%s" in %s`+"\n", quote.Quote, quote.Anime)
	}
}

func fetchQuoteByAnimeName(name string, page string) {
	var url string = fmt.Sprintf("https://animechan.vercel.app/api/quotes/anime?title=%s&page=%s", name, page)
	var quotes []*quote

	body := responseData(url)

	err := json.Unmarshal(body, &quotes)
	if err != nil {
		log.Panicf("error when unmarshaling data random quote: %v\n", err)
	}

	for _, quote := range quotes {
		fmt.Printf(`"%s" by %s`+"\n", quote.Quote, quote.Character)
	}

}

func fetchRandomQuote() {
	var url string = "https://animechan.vercel.app/api/random"
	var quote *quote

	body := responseData(url)

	err := json.Unmarshal(body, &quote)
	if err != nil {
		log.Panicf("error when unmarshaling data random quote: %v\n", err)
	}

	fmt.Printf("%s\n \t %s in %s\n", quote.Quote, quote.Character, quote.Anime)
}

func quoteAnimeCmdExecute(args []string) {
	if len(args) < 1 {
		fetchRandomQuote()
		return
	}

	if len(args) == 1 {
		title := args[0]

		fetchQuoteByAnimeName(title, "1")
		return
	}

	if len(args) > 1 {
		title := args[0]
		page := args[1]

		fetchQuoteByAnimeName(title, page)
		return
	}
}
