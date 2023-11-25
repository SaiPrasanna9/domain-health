/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check command checks and gets domain status",
	Long:  `check command is a command in domain-health cli-tool that checks if a given domain is reachable or not and outputs the result as status.`,
	Run: func(cmd *cobra.Command, args []string) {
		domainFlag, _ := cmd.Flags().GetString("domain")
		portFlag, _ := cmd.Flags().GetString("port")
		if portFlag == "" {
			portFlag = "80"
		}

		status := Check(domainFlag, portFlag)
		fmt.Print(status)
	},
}

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)
	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable \n FROM: %v \n TO: %v \n", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status
}
func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	checkCmd.Flags().String("domain", "", "provide domain url")
	checkCmd.Flags().String("port", "", "provide domain port")
}
