package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/m1/go-generate-password/generator"
	"github.com/spf13/cobra"
)

// generatePasswordCmd represents the generatePassword command
var generatePasswordCmd = &cobra.Command{
	Use:   "generate",
	Short: "A command generate random password",
	Long:  `Generate random password to prevent lacking`,
	Run: func(cmd *cobra.Command, args []string) {
		generatePasswordExecute()
	},
}

func init() {
	rootCmd.AddCommand(generatePasswordCmd)
}

func generateRandomPassword() string {
	config := generator.Config{
		Length:                     20,
		IncludeSymbols:             true,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}
	g, _ := generator.New(&config)

	pwd, err := g.Generate()
	if err != nil {
		log.Fatalf("Error when generate password: %v", err)
	}

	return *pwd
}

func generatePasswordExecute() {
	password := generateRandomPassword()
	fmt.Printf("Your password: %v\n", color.RedString(password))
}
