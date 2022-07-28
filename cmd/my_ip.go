package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var myIpCmd = &cobra.Command{
	Use:   "my_ip",
	Short: "Get current IP address of current device",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		myIpCmdExecute(cmd, args)
	},
}

var myIpError = func(err string) {
	log.Fatalf("myIpCmdExecute - %v\n", err)
}

func init() {
	rootCmd.AddCommand(myIpCmd)

	// sub-command
	myIpCmd.PersistentFlags().String("public", "", "Get public IP in global")
	myIpCmd.PersistentFlags().String("local", "", "Get local IP")
}

func fetchCurrentIpLocal() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("myIp - myIpCmdExecute - Error: %s", err)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Printf("IP Address: %s\n", color.GreenString(ipnet.IP.String()))
			}
		}
	}
}

func fetchCurrentIpGlobal() {
	url := "https://api.ipify.org?format=text"
	data := ResponseData(url)

	fmt.Printf("IP Address: %s\n", color.GreenString(string(data)))
}

func myIpCmdExecute(cmd *cobra.Command, args []string) {
	local, err := cmd.Flags().GetString("local")
	if err != nil {
		myIpError(fmt.Sprintf("Error: local %s", err))
	}

	public, err := cmd.Flags().GetString("public")
	if err != nil {
		myIpError(fmt.Sprintf("Error: public %s", err))
	}

	if local == "" && public == "" {
		fetchCurrentIpLocal()
	}

	if local != "" {
		fetchCurrentIpLocal()
	}

	if public != "" {
		fetchCurrentIpGlobal()
	}
}
