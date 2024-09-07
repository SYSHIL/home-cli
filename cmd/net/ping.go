/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"net/http"

	"github.com/spf13/cobra"
)

var (
	url string
)

func ping(url string) (int, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping is a command that sends a ICMP echo request to a host",
	Long:  `Ping is a command that sends a ICMP echo request to a host.`,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(url); err != nil {
			cmd.PrintErrln(err)
		} else {
			cmd.Println(resp)
		}
	},
}

func init() {

	pingCmd.Flags().StringVarP(&url, "url", "u", "localhost", "The url to ping")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	NetCmd.AddCommand(pingCmd)

}
