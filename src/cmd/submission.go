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
		const sheetName = "スキルシート"
		const worksheetSize = 9 // A4 210 x 297 mm

		f := excelize.NewFile()
		f.SetSheetName("Sheet1", sheetName)
		f.SetPageLayout(sheetName, excelize.PageLayoutPaperSize(worksheetSize))

		set := func(position string, val interface{}) {
			if err := f.SetCellValue(sheetName, position, val); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		merge := func(from, to string) {
			if err := f.MergeCell(sheetName, from, to); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		set("A1", "スキルシート")
		merge("A1", "F1")

		set("A2", fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")))
		merge("A2", "F2")

		set("A3", "フリガナ")
		set("B3", cfg.Section("").Key("kana").String())

		set("A4", "名前")
		set("B4", cfg.Section("").Key("name").String())

		set("C3", "ニックネーム")
		set("C4", *attribute.Nickname)

		set("D3", "年齢")
		birthYear := *attribute.Birthday.Year
		birthMonth := *attribute.Birthday.Month
		birthDay := *attribute.Birthday.Day
		set("D4", age(birthYear, birthMonth, birthDay, time.Now()))

		set("E3", "メールアドレス")
		set("E4", cfg.Section("").Key("mail").String())

		set("F3", "Gravatar")

		if err := f.SaveAs("skill_sheet.xlsx"); err != nil {
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
