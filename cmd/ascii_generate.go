package cmd

import (
	"fmt"
	"log"

	_ "image/jpeg"

	"github.com/spf13/cobra"
	asciicanvas "github.com/tompng/go-ascii-canvas"
)

var asciiGenerateCmd = &cobra.Command{
	Use:   "ascii",
	Short: "A command generate ascii image with Go",
	Long:  `Generate ascii image with Go`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := args[0]
		buffer := asciiGenerate(fileName)
		fmt.Printf("%v \n", buffer)
	},
}

func init() {
	rootCmd.AddCommand(asciiGenerateCmd)
}

func asciiGenerate(fileName string) *asciicanvas.ImageBuffer {
	image, err := asciicanvas.NewImageBufferFromFile(fileName)
	if err != nil {
		log.Println("Cannot read file")
	}

	canvas := asciicanvas.NewImageBuffer(120, 90)
	canvas.Draw(image, 0, 0, 80, 80)
	return canvas
}
