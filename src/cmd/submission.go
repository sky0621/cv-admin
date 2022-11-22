/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gopkg.in/ini.v1"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/xuri/excelize/v2"

	"github.com/sky0621/cv-admin/src/rest"
)

// submissionCmd represents the submission command
var submissionCmd = &cobra.Command{
	Use:   "submission",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		/*
		 * INIから個人情報を読み込み
		 */
		cfg, err := ini.Load(".private.ini")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		/*
		 * SQLiteからキャリア情報を読み込み（API経由）
		 */
		cli := http.DefaultClient

		attribute, err := requestUserInfo(cli, fmt.Sprintf("users/%d/attribute", targetUserID), &rest.UserAttribute{})
		if err != nil {
			os.Exit(1)
		}

		/*
		 * Excelに書き込み
		 */
		const sheetName = "キャリアシート"

		f := excelize.NewFile()
		f.SetSheetName("Sheet1", sheetName)

		set := func(position string, val interface{}) {
			if err := f.SetCellValue(sheetName, position, val); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		set("B2", "スキルシート")

		set("B4", fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")))

		set("B5", "フリガナ")
		set("C5", cfg.Section("").Key("kana").String())

		set("B6", "名前")
		set("C6", cfg.Section("").Key("name").String())

		set("D5", "ニックネーム")
		set("D6", *attribute.Nickname)

		set("E5", "年齢")
		birthYear := *attribute.Birthday.Year
		birthMonth := *attribute.Birthday.Month
		birthDay := *attribute.Birthday.Day
		set("E6", age(birthYear, birthMonth, birthDay, time.Now()))

		if err := f.SaveAs("career_sheet.xlsx"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(submissionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submissionCmd.PersistentFlags().String("foo", "", "A help for foo")
	submissionCmd.PersistentFlags().IntVarP(&targetUserID, "userid", "u", 29, "target user id")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submissionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func age(birthYear, birthMonth, birthDay int, now time.Time) int {
	age := now.Year() - birthYear

	if int(now.Month()) > birthMonth {
		return age
	}

	if int(now.Month()) == birthMonth && now.Day() >= birthDay {
		return age
	}

	return age - 1
}
