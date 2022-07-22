package cmd

import (
	"encoding/json"
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
	Use:   "anime [OPTIONS] [PAGE]",
	Short: "Fetch quote anime",
	Long:  `Use to fetch anime's quote: random, with charactor, with anime`,
	Run: func(cmd *cobra.Command, args []string) {
		quoteAnimeCmdExecute(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(quoteAnimeCmd)

	quoteAnimeCmd.PersistentFlags().String("name", "", "Search quote by name of anime")
	quoteAnimeCmd.PersistentFlags().String("charactor", "", "Search quote by name of charactor")
}

func fetchQuoteByCharactorName(name string, page string) {
	var url string = fmt.Sprintf("https://animechan.vercel.app/api/quotes/character?name=%s&page=%s", name, page)
	var quotes []*quote

	body := ResponseData(url)

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

	body := ResponseData(url)

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

	body := ResponseData(url)

	err := json.Unmarshal(body, &quote)
	if err != nil {
		log.Panicf("error when unmarshaling data random quote: %v\n", err)
	}

	fmt.Printf("%s\n", color.YellowString(quote.Quote))
	fmt.Printf("\t %s in %s\n", color.RedString(quote.Character), color.BlueString(quote.Anime))
}

func quoteAnimeCmdExecute(cmd *cobra.Command, args []string) {
	animeName, err := cmd.Flags().GetString("name")
	if err != nil {
		log.Fatalf("Error with use flags quoteAnimeExecute: %v", err)
	}

	charactor, err := cmd.Flags().GetString("charactor")
	if err != nil {
		log.Fatalf("Error with use flags quoteAnimeExecute: %v", err)
	}

	if animeName == "" && charactor == "" {
		fetchRandomQuote()
	}

	var page string = "1"

	if len(args) > 0 {
		page = args[0]
	}

	if animeName != "" {
		fetchQuoteByAnimeName(animeName, page)
	}

	if charactor != "" {
		fetchQuoteByCharactorName(charactor, page)
	}

}
