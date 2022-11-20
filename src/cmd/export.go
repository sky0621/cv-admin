/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/sky0621/cv-admin/src/rest"
)

// 厳密にはプロトコル＋FQDNだけど。
var requestFQDN = "http://localhost:8080"

var targetUserID int

var exportDir string

func exportJSON[T any](cli *http.Client, urlPath, outputFileName string, str T) bool {
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

	err = json.Unmarshal(resBody, str)
	if err != nil {
		fmt.Printf("[%s][Unmarshal] %v\n", urlPath, err)
		return false
	}

	jsonBytes, err := json.MarshalIndent(str, "", "  ")
	if err != nil {
		fmt.Printf("[%s][MarshalIndent] %v\n", urlPath, err)
		return false
	}

	jsonFile, err := os.Create(fmt.Sprintf("%s/%s.json", exportDir, outputFileName))
	if err != nil {
		fmt.Printf("[%s][Create json] %v\n", urlPath, err)
		return false
	}
	defer jsonFile.Close()

	_, err = io.WriteString(jsonFile, string(jsonBytes))
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

		if !exportJSON(cli, fmt.Sprintf("users/%d/attribute", targetUserID), "attribute", &rest.UserAttribute{}) {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%d/activities", targetUserID), "activities", &[]*rest.UserActivity{}) {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%d/qualifications", targetUserID), "qualifications", &[]*rest.UserQualification{}) {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%d/careergroups", targetUserID), "careergroups", &[]*rest.UserCareerGroup{}) {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%d/notes", targetUserID), "notes", &[]*rest.UserNote{}) {
			os.Exit(1)
		}

		if !exportJSON(cli, fmt.Sprintf("users/%d/skills", targetUserID), "skills", &[]*rest.UserSkillTag{}) {
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
	exportCmd.PersistentFlags().IntVarP(&targetUserID, "userid", "u", 29, "target user id")
	exportCmd.PersistentFlags().StringVarP(&exportDir, "dir", "d", "/tmp", "export directory")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
