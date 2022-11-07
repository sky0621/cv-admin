/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// 厳密にはプロトコル＋FQDNだけど。
var requestFQDN = "http://localhost:8080"

var targetUserID = "26"

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli := http.DefaultClient
		/*
		 * user attribute
		 */
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/users/%s/attribute", requestFQDN, targetUserID), nil)
		if err != nil {
			fmt.Printf("[user attribute][NewRequest] %v\n", err)
			os.Exit(1)
		}

		res, err := cli.Do(req)
		if err != nil {
			fmt.Printf("[user attribute][Do] %v\n", err)
			os.Exit(1)
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("[user attribute][ReadAll] %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", resBody)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
