package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type quote struct {
	Anime     string `json:"anime" example:"anime"`
	Character string `json:"character" example:"character"`
	Quote     string `json:"quote" example:"quote"`
}

var quoteAnimeCmd = &cobra.Command{
	Use:   "quote",
	Short: "Fetch quote anime",
	Long:  `Use to fetch anime's quote: random, with charactor, with anime`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

var quoteAnimeRandomCmd = &cobra.Command{
	Use:   "random",
	Short: "Fetch quote anime",
	Long:  `Use to fetch anime's quote: random, with charactor, with anime`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchRandomQuote()
	},
}

var quoteByNameAnimeCmd = &cobra.Command{
	Use:   "anime",
	Short: "Fetch quote anime",
	Long:  `Use to fetch anime's quote: random, with charactor, with anime`,
	Run: func(cmd *cobra.Command, args []string) {
		fetchQuoteByAnimeName(name, page)
	},
}

var quoteByNameCharacterCmd = &cobra.Command{
	Use:   "charactor",
	Short: "Fetch quote anime",
	Run: func(cmd *cobra.Command, args []string) {
		fetchQuoteByCharactorName(name, page)
	},
}

func init() {
	rootCmd.AddCommand(quoteAnimeCmd)

	// sub-command
	quoteAnimeCmd.AddCommand(quoteAnimeRandomCmd)
	quoteAnimeCmd.AddCommand(quoteByNameAnimeCmd)
	quoteAnimeCmd.AddCommand(quoteByNameCharacterCmd)

	// binding flag
	quoteAnimeCmd.PersistentFlags().StringVar(&name, "name", "", "Search quote by name of anime")
	quoteAnimeCmd.PersistentFlags().StringVar(&page, "page", "1", "Search quote by name of charactor")
}

func fetchQuoteByCharactorName(name string, page string) {
	var url string = fmt.Sprintf("https://animechan.vercel.app/api/quotes/character?name=%s&page=%s", name, page)
	var quotes []*quote

	body := ResponseData(url)

	err := json.Unmarshal(body, &quotes)
	if err != nil {
		log.Panicf("quote_anime - fetchQuoteByCharacterName - Error: %v\n", err)
	}

	for _, quote := range quotes {
		fmt.Printf(`"%s" in %s`+"\n", quote.Quote, quote.Anime)
	}
}

func fetchQuoteByAnimeName(name string, page string) {
	var url string = fmt.Sprintf("https://animechan.vercel.app/api/quotes/anime?title=%s&page=%s", name, page)
	var quotes []*quote

	body := ResponseData(url)

	err := json.Unmarshal(body, &quotes)
	if err != nil {
		log.Panicf("quote_anime - fetchQuoteByAnimeName - Error: %v\n", err)
	}

	for _, quote := range quotes {
		fmt.Printf(`"%s" by %s`+"\n", quote.Quote, quote.Character)
	}
}

func fetchRandomQuote() {
	var url string = "https://animechan.vercel.app/api/random"
	var quote *quote

	body := ResponseData(url)

	err := json.Unmarshal(body, &quote)
	if err != nil {
		log.Panicf("quote_anime - fetchRandomQuote - Error: %v\n", err)
	}

	fmt.Printf("%s\n", color.YellowString(quote.Quote))
	fmt.Printf("\t %s in %s\n", color.RedString(quote.Character), color.BlueString(quote.Anime))
}
