package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerBackupCmd represents the dockerBackup command
var dockerBackupCmd = &cobra.Command{
	Use:   "docker-backup",
	Short: "Command docker backup",
	Long:  `Command docker backup`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dockerBackup called")
	},
}

func init() {
	rootCmd.AddCommand(dockerBackupCmd)
}
