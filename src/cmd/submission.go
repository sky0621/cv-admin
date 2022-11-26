/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	s "github.com/sky0621/cv-admin/src/submission"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sky0621/cv-admin/src/rest"
	"github.com/spf13/cobra"
)

func requestAvatarImage(cli *http.Client, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("[%s][NewRequest] %v\n", url, err)
		return nil, err
	}

	res, err := cli.Do(req)
	if err != nil {
		fmt.Printf("[%s][Do] %v\n", url, err)
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[%s][ReadAll] %v\n", url, err)
		return nil, err
	}

	return resBody, nil
}

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
		/* **********************************************************
		 * INIから個人情報を読み込み
		 */
		cfg := s.NewConfig()

		/* **********************************************************
		 * SQLiteからキャリア情報を読み込み（API経由）
		 */
		cli := http.DefaultClient

		attribute, err := requestUserInfo(cli, fmt.Sprintf("users/%d/attribute", targetUserID), &rest.UserAttribute{})
		if err != nil {
			os.Exit(1)
		}

		var avatarImgBytes []byte
		if attribute != nil {
			avatarImgBytes, err = requestAvatarImage(cli, *attribute.AvatarUrl)
			if err != nil {
				os.Exit(1)
			}
		}

		qualifications, err := requestUserInfo(cli, fmt.Sprintf("users/%d/qualifications", targetUserID), &[]rest.UserQualification{})
		if err != nil {
			os.Exit(1)
		}

		activities, err := requestUserInfo(cli, fmt.Sprintf("users/%d/activities", targetUserID), &[]rest.UserActivity{})
		if err != nil {
			os.Exit(1)
		}

		skillTags, err := requestUserInfo(cli, fmt.Sprintf("users/%d/skills", targetUserID), &[]rest.UserSkillTag{})
		if err != nil {
			os.Exit(1)
		}

		/* **********************************************************
		 * Excelに書き込み
		 */
		w := s.NewExcelizeWrapper(
			s.SheetName("スキルシート"),
			// MEMO: パスワードロックしても別シート作ってコピペはできるので（誤操作を防ぎたいわけではないので）意味なし。。
			// s.SheetPassword(s.GetConfigValue(cfg, s.Password)),
			s.PaperSize(9), // A4 210 x 297 mm
		)

		/*
		 * 列ごとの設定
		 */
		{
			w.Width("A", "A", 15)
			w.Width("B", "B", 30)
			w.Width("C", "C", 20)
			w.Width("D", "D", 10)
			w.Width("E", "E", 40)
			w.Width("F", "F", 15)
		}

		/*
		 * １行目
		 */
		{
			w.Height(1, 30)
			w.Set("A1", "スキルシート")
			w.Merge("A1", "F1")
			w.CellStyle("A1", s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.Border),
				s.Font(s.SheetTitleFont),
			))
		}

		/*
		 * ２行目
		 */
		{
			w.Height(2, 20)
			w.Set("A2", fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")))
			w.Merge("A2", "F2")
			w.CellStyle("A2", s.NewStyle(
				s.Alignment(s.HRightAlignment),
				s.Font(s.LastUpdatedFont),
			))
		}

		/*
		 * ３行目
		 */
		{
			w.Height(3, s.RowBaseHeight)
			w.Set("A3", "フリガナ")
			w.HeaderCellStyle("A3")
			w.Set("B3", s.GetConfigValue(cfg, s.Kana))
			w.CellStyle("B3", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Set("C3", "ニックネーム")
			w.Set("D3", "年齢")
			w.Set("E3", "メールアドレス")
			w.Set("F3", "Gravatar")
			w.HeaderCellRangeStyle("C3", "F3")
			w.Merge("F4", "F6")
		}

		/*
		 * ４行目
		 */
		{
			w.Height(4, 45)
			w.Set("A4", "名前")
			w.HeaderCellStyle("A4")
			w.Set("B4", cfg.Section("").Key("name").String())
			w.CellStyle("B4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
				s.Font(s.NameFont),
			))

			// ニックネーム
			w.Set("C4", *attribute.Nickname)
			w.CellStyle("C4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))

			// 年齢
			bDay := *attribute.Birthday
			w.Set("D4", age(*bDay.Year, *bDay.Month, *bDay.Day, time.Now()))
			w.CellStyle("D4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))

			// メールアドレス
			w.Set("E4", s.GetConfigValue(cfg, s.Mail))
			w.CellStyle("E4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))

			// Gravatar
			if avatarImgBytes != nil {
				w.AddPictureFromBytes("F4", "avatar", ".png", avatarImgBytes)
			}
			w.CellStyle("F4", s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.Border),
			))
		}

		/*
		 * 5行目
		 */
		{
			w.Height(5, s.RowBaseHeight)
			w.Set("A5", "居住地")
			w.HeaderCellStyle("A5")
			w.Set("B5", s.GetConfigValue(cfg, s.CityOfResidence))
			w.CellStyle("B5", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Set("C5", "最寄り駅")
			w.HeaderCellStyle("C5")
			w.Set("D5", s.GetConfigValue(cfg, s.NearestStation))
			w.CellStyle("D5", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Merge("D5", "E5")
		}

		/*
		 * 6行目
		 */
		{
			w.Height(6, s.RowBaseHeight)
			w.Set("A6", "最終学歴")
			w.HeaderCellStyle("A6")
			w.Set("B6", s.GetConfigValue(cfg, s.EducationalBackground))
			w.CellStyle("B6", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Merge("B6", "E6")
		}

		/*
		 * 8行目
		 */
		{
			w.Height(8, s.RowBaseHeight)
			w.Set("A8", "資格")
			w.Merge("A8", "B8")
			w.Set("C8", "取得年月日")
			w.Set("D8", "URL")
			w.Merge("D8", "F8")
			w.HeaderCellRangeStyle("A8", "F8")
		}

		/*
		 * 9行目〜
		 * 資格情報の件数分
		 */
		rowNo := 8
		{
			if qualifications != nil {
				for _, q := range *qualifications {
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight)

					// 資格
					w.Set(s.Cell("A", rowNo), s.QualificationName(q.Organization, q.Name))
					w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("B", rowNo))

					// 取得年月日
					w.Set(s.Cell("C", rowNo), s.QualificationGotDate(q.GotDate))
					w.CellStyle(s.Cell("C", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))

					// URL
					w.Set(s.Cell("D", rowNo), s.URL(q.Url))
					w.CellExternalHyperLink(s.Cell("D", rowNo), s.URL(q.Url))
					w.CellStyle(s.Cell("D", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))

					w.Merge(s.Cell("D", rowNo), s.Cell("F", rowNo))
				}
			}
		}

		/*
		 * アウトプットラベル
		 */
		rowNo += 2
		{
			w.Height(rowNo, s.RowBaseHeight)
			w.Set(s.Cell("A", rowNo), "アウトプット")
			w.Merge(s.Cell("A", rowNo), s.Cell("C", rowNo))
			w.Set(s.Cell("D", rowNo), "URL")
			w.Merge(s.Cell("D", rowNo), s.Cell("F", rowNo))
			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("F", rowNo))
		}

		/*
		 * アウトプット
		 */
		{
			if activities != nil {
				for _, a := range *activities {
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight)

					// アウトプット
					w.Set(s.Cell("A", rowNo), *a.Name)
					w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("C", rowNo))

					// URL
					w.Set(s.Cell("D", rowNo), s.URL(a.Url))
					w.CellExternalHyperLink(s.Cell("D", rowNo), s.URL(a.Url))
					w.CellStyle(s.Cell("D", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))

					w.Merge(s.Cell("D", rowNo), s.Cell("F", rowNo))
				}
			}
		}

		/*
		 * アウトプット（from .private.ini）
		 */
		{
			rowNo += 1
			w.Height(rowNo, s.RowBaseHeight)

			// アウトプット
			w.Set(s.Cell("A", rowNo), "スキルシート(Web版)")
			w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Merge(s.Cell("A", rowNo), s.Cell("C", rowNo))

			// URL
			w.Set(s.Cell("D", rowNo), s.GetConfigValue(cfg, s.CvWeb))
			w.CellExternalHyperLink(s.Cell("D", rowNo), s.GetConfigValue(cfg, s.CvWeb))
			w.CellStyle(s.Cell("D", rowNo), s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))

			w.Merge(s.Cell("D", rowNo), s.Cell("F", rowNo))
		}

		/*
		 * PRラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)
			w.Set(s.Cell("A", rowNo), "PR")
			w.Merge(s.Cell("A", rowNo), s.Cell("F", rowNo))
			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("F", rowNo))
		}

		/*
		 * PR
		 */
		{
			rowNo += 1
			w.Height(rowNo, s.RowMaxHeight) // TODO: PR行数に応じて（高さが収まらない場合があるので）セルを分ける対応が必要
			w.Set(s.Cell("A", rowNo), *attribute.Pr)
			w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.Border),
			))
			w.Merge(s.Cell("A", rowNo), s.Cell("F", rowNo))
		}

		/*
		 * スキルラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)
			w.Set(s.Cell("A", rowNo), "スキル")
			w.Merge(s.Cell("A", rowNo), s.Cell("F", rowNo))
			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("F", rowNo))
		}

		/*
		 * スキル
		 */
		{
			if skillTags != nil {
				for _, t := range *skillTags {
					/*
					 * スキルカテゴリ
					 */
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight)

					w.Set(s.Cell("A", rowNo), *t.TagName)
					w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.Border),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("F", rowNo))

					if t.Skills != nil {
						for _, s := range *t.Skills {
							// FIXME:
							fmt.Println(*s.Name)
						}
					}
				}
			}
		}

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
	// sCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
