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

var findPublicIp = &cobra.Command{
	Use:   "parse", //
	Short: "Parse DNS to Ip address",
	Run: func(cmd *cobra.Command, args []string) {
		parseDnsToIpAddress(dns)
	},
}

func init() {
	rootCmd.AddCommand(myIpCmd)

	// sub-command
	myIpCmd.AddCommand(myIpLocalCmd)
	myIpCmd.AddCommand(myIpPublicCmd)
	myIpCmd.AddCommand(findPublicIp)

	myIpCmd.PersistentFlags().StringVar(&dns, "dns", "", "domain name server")
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

func parseDnsToIpAddress(dnsName string) {
	ips, _ := net.LookupIP(dnsName)

	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			log.Printf("IPv4: %s\n", color.GreenString(ipv4.String()))
		}
	}
}
