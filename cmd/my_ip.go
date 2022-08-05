package cmd

import (
	"log"
	"net"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var myIpError = func(err string) {
	log.Fatalf("myIpCmdExecute - %v\n", err)
}

var myIpCmd = &cobra.Command{
	Use:   "ip",
	Short: "Get current IP address of current device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var myIpLocalCmd = &cobra.Command{
	Use:   "local",
	Short: "Get current IP in local",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fetchCurrentIpLocal()
	},
}

var myIpPublicCmd = &cobra.Command{
	Use:   "public",
	Short: "Get public IP internet ",
	Run: func(cmd *cobra.Command, args []string) {
		fetchCurrentIpGlobal()
	},
}

func init() {
	rootCmd.AddCommand(myIpCmd)

	// sub-command
	myIpCmd.AddCommand(myIpLocalCmd)
	myIpCmd.AddCommand(myIpPublicCmd)
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

	log.Printf("IP Address: %s\n", color.GreenString(string(data)))
}
