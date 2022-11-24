/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/sky0621/cv-admin/src/submission"
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
		const worksheetSize = 9 // A4 210 x 297 mm
		w := submission.NewExcelizeWrapper("スキルシート", worksheetSize)

		w.Width("A", "A", 10)
		w.Width("B", "B", 20)
		w.Width("C", "C", 15)
		w.Width("D", "D", 5)
		w.Width("E", "E", 40)
		w.Width("F", "F", 10)

		/*
		 * １行目
		 */
		w.Height(1, 30)
		w.Merge("A1", "F1")
		w.Text("A1", []excelize.RichTextRun{
			{
				Text: "スキルシート",
				Font: &excelize.Font{
					Size: 16,
					Bold: true,
				},
			},
		})
		w.CellStyle("A1", "A1", w.CenterStyleID())

		/*
		 * ２行目
		 */
		w.Height(2, 20)
		w.Merge("A2", "F2")
		w.Text("A2", []excelize.RichTextRun{
			{
				Text: fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")),
				Font: &excelize.Font{
					Size: 9,
				},
			},
		})
		w.CellStyle("A2", "A2", w.RightStyleID())

		/*
		 * ３行目
		 */
		w.Height(3, 25)
		w.Set("A3", "フリガナ")
		w.Set("B3", cfg.Section("").Key("kana").String())
		w.Set("C3", "ニックネーム")
		w.Set("D3", "年齢")
		w.Set("E3", "メールアドレス")
		w.Set("F3", "Gravatar")

		/*
		 * ４行目
		 */
		w.Height(4, 45)
		w.Set("A4", "名前")
		w.Text("B4", []excelize.RichTextRun{
			{
				Text: cfg.Section("").Key("name").String(),
				Font: &excelize.Font{
					Size: 14,
					Bold: true,
				},
			},
		})
		w.Set("C4", *attribute.Nickname)

		birthYear := *attribute.Birthday.Year
		birthMonth := *attribute.Birthday.Month
		birthDay := *attribute.Birthday.Day
		w.Set("D4", age(birthYear, birthMonth, birthDay, time.Now()))

		w.Set("E4", cfg.Section("").Key("mail").String())

		w.SaveAs("skill_sheet.xlsx")
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
