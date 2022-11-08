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

var outputDir = "/tmp/outputjson"

func exportJSON(cli *http.Client, urlPath, outputFileName string) bool {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", requestFQDN, urlPath), nil)
	if err != nil {
		fmt.Printf("[%s][NewRequest] %v\n", urlPath, err)
		return false
	}

	res, err := cli.Do(req)
	if err != nil {
		fmt.Printf("[%s][Do] %v\n", urlPath, err)
		return false
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[%s][ReadAll] %v\n", urlPath, err)
		return false
	}

	attribute, err := os.Create(fmt.Sprintf("%s/%s.json", outputDir, outputFileName))
	if err != nil {
		fmt.Printf("[%s][Create json] %v\n", urlPath, err)
		return false
	}
	defer attribute.Close()

	_, err = io.WriteString(attribute, string(resBody))
	if err != nil {
		fmt.Printf("[%s][Write json] %v\n", urlPath, err)
		return false
	}

	return true
}

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

		if !exportJSON(cli, fmt.Sprintf("users/%s/attribute", targetUserID), "attribute") {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%s/activities", targetUserID), "activities") {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%s/qualifications", targetUserID), "qualifications") {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%s/careergroups", targetUserID), "careergroups") {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%s/notes", targetUserID), "notes") {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%s/skills", targetUserID), "skills") {
			os.Exit(1)
		}
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
