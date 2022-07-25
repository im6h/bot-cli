package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var copyEnvCmd = &cobra.Command{
	Use:   "cp_env",
	Short: "A command to make .env.example from .env",
	Long:  `Use command to make .env.example from .env to public repo`,
	Run: func(cmd *cobra.Command, args []string) {
		CopyEnvExecute()
	},
}

var currentPath string
var envPath string
var envDest string

func init() {
	rootCmd.AddCommand(copyEnvCmd)
	currentPath, _ = os.Getwd()
	envPath = fmt.Sprintf("%s/%s", currentPath, ".env")
	envDest = fmt.Sprintf("%s/%s", currentPath, ".env.example")
}

// TODO: check file .env exist in current path
func checkEnvExist() (err error) {

	if _, err := os.Stat(envPath); err != nil {
		return fmt.Errorf("checkEnvExist - Error: %s", err)
	}

	return nil
}

// TODO: execute function copy
func copyLineByLine(file *os.File, str string) {
	regex, err := regexp.Compile("/.*=/gm")
	if err != nil {

	}

	name := regex.Split(str, 2)

	if _, err := file.WriteString(fmt.Sprintln(name[0])); err != nil {
		log.Panicf("CopyLineByLine - WriteString - Error: %s", err)
	}
}

func CopyEnvExecute() {
	isExist := checkEnvExist()

	if isExist != nil {
		log.Fatalf(color.RedString("CopyEnvExecute - : %s\n", isExist))
	}

	file, err := os.Open(envPath)
	if err != nil {
		log.Fatalf(color.RedString("CopyEnvExecute - Open - Error: %s\n", err))
	}

	destFile, err := os.OpenFile(envDest, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(color.RedString("CopyEnvExecute - Open - Error: %s\n", err))
	}

	defer file.Close()
	defer destFile.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		dataInLine := sc.Text()
		copyLineByLine(destFile, dataInLine)
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("CopyEnvExecute - Scan Line - Error: %v", err)
		return
	}
}
