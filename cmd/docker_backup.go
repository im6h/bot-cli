package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// dockerBackupCmd represents the dockerBackup command
var dockerBackupCmd = &cobra.Command{
	Use:   "docker-backup",
	Short: "Command docker backup",
	Long:  `Command docker backup`,
	Run: func(cmd *cobra.Command, args []string) {
		db := args[0]
		containerId := args[1]
		dbPassword := args[2]
		username := args[3]
		databaseName := args[4]
		if containerId == "" {
			log.Fatal("Need container id")
		}
		if dbPassword == "" {
			log.Fatal("Need dbPassword")
		}
		if username == "" {
			log.Fatal("Need username")
		}
		if databaseName == "" {
			log.Fatal("Need databaseName")
		}

		switch db {
		case "pg":
			{
				backupPg(containerId, dbPassword, username, databaseName)
				break
			}
		case "mysql":
			{
				break
			}
		default:
			{
				log.Fatal("Need database type")
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dockerBackupCmd)
}

func backupPg(containerId string, dbPassword string, username string, databaseName string) {
	cmdString := fmt.Sprintf(`exec -i %s /bin/bash -c "PGPASSWORD=%s pg_dump --username %s %s" > /Users/oraichain/Desktop/work/database-sql/testnet/dump_date.sql`, containerId, dbPassword, username, databaseName)

	cmd := exec.Command("docker", cmdString)

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
