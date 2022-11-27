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
	"strings"
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
			s.SheetView(),
		)

		/*
		 * 列ごとの設定
		 */
		{
			w.Width("A", "Z", 5)
			w.Width("AA", "AA", 0.5)
		}

		/*
		 * １行目
		 */
		{
			w.Height(1, 30)
			w.Set("A1", "スキルシート")
			w.Merge("A1", "Z1")
			w.CellStyle("A1", s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
				s.Font(s.SheetTitleFont),
			))
			// 枠線「右」が機能していないための措置
			w.CellStyle("AA1", s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * ２行目
		 */
		{
			w.Height(2, 20)
			w.Set("A2", fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")))
			w.Merge("A2", "Z2")
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
			w.Merge("A3", "C3")

			w.Set("D3", s.GetConfigValue(cfg, s.Kana))
			w.CellStyle("D3", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("D3", "I3")

			w.Set("J3", "ニックネーム")
			w.Merge("J3", "M3")

			w.Set("N3", "年齢")
			w.Merge("N3", "O3")

			w.Set("P3", "メールアドレス")
			w.Merge("P3", "W3")

			w.Set("X3", "Gravatar")
			w.Merge("X3", "Z3")

			w.HeaderCellRangeStyle("J3", "Z3")

			w.Merge("X4", "Z4")
			w.Merge("X5", "Z5")
			w.Merge("X6", "Z6")
			w.Merge("X4", "X6")

			// 枠線「右」が機能していないための措置
			w.CellStyle("AA3", s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * ４行目
		 */
		{
			w.Height(4, 45)
			w.Set("A4", "名前")
			w.HeaderCellStyle("A4")
			w.Merge("A4", "C4")

			w.Set("D4", cfg.Section("").Key("name").String())
			w.CellStyle("D4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
				s.Font(s.NameFont),
			))
			w.Merge("D4", "I4")

			// ニックネーム
			w.Set("J4", *attribute.Nickname)
			w.CellStyle("J4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("J4", "M4")

			// 年齢
			bDay := *attribute.Birthday
			w.Set("N4", age(*bDay.Year, *bDay.Month, *bDay.Day, time.Now()))
			w.CellStyle("N4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("N4", "O4")

			// メールアドレス
			w.Set("P4", s.GetConfigValue(cfg, s.Mail))
			w.CellStyle("P4", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("P4", "W4")

			// Gravatar
			if avatarImgBytes != nil {
				w.AddPictureFromBytes("X4", "avatar", ".png", avatarImgBytes)
			}
			w.CellStyle("X4", s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
			))

			// 枠線「右」が機能していないための措置
			w.CellStyle("AA4", s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 5行目
		 */
		{
			w.Height(5, s.RowBaseHeight)
			w.Set("A5", "居住地")
			w.HeaderCellStyle("A5")
			w.Merge("A5", "C5")

			w.Set("D5", s.GetConfigValue(cfg, s.CityOfResidence))
			w.CellStyle("D5", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("D5", "I5")

			w.Set("J5", "最寄り駅")
			w.HeaderCellStyle("J5")
			w.Merge("J5", "M5")

			w.Set("N5", s.GetConfigValue(cfg, s.NearestStation))
			w.CellStyle("N5", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("N5", "W5")

			// 枠線「右」が機能していないための措置
			w.CellStyle("AA5", s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 6行目
		 */
		{
			w.Height(6, s.RowBaseHeight)
			w.Set("A6", "最終学歴")
			w.HeaderCellStyle("A6")
			w.Merge("A6", "C6")

			w.Set("D6", s.GetConfigValue(cfg, s.EducationalBackground))
			w.CellStyle("D6", s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge("D6", "W6")

			// 枠線「右」「下」が機能していないための措置
			w.CellStyle("X6", s.NewStyle(s.Borders(s.BottomBorder)))
			w.CellStyle("AA6", s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 8行目
		 */
		{
			w.Height(8, s.RowBaseHeight)
			w.Set("A8", "資格")
			w.Merge("A8", "I8")

			w.Set("J8", "取得年月日")
			w.Merge("J8", "M8")

			w.Set("N8", "URL")
			w.Merge("N8", "Z8")

			w.HeaderCellRangeStyle("A8", "Z8")

			// 枠線「右」が機能していないための措置
			w.CellStyle("AA8", s.NewStyle(s.Borders(s.LeftBorder)))
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
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("I", rowNo))

					// 取得年月日
					w.Set(s.Cell("J", rowNo), s.QualificationGotDate(q.GotDate))
					w.CellStyle(s.Cell("J", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("J", rowNo), s.Cell("M", rowNo))

					// URL
					w.Set(s.Cell("N", rowNo), s.URL(q.Url))
					w.CellExternalHyperLink(s.Cell("N", rowNo), s.URL(q.Url))
					w.CellStyle(s.Cell("N", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("N", rowNo), s.Cell("Z", rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
			w.Merge(s.Cell("A", rowNo), s.Cell("I", rowNo))

			w.Set(s.Cell("J", rowNo), "URL")
			w.Merge(s.Cell("J", rowNo), s.Cell("Z", rowNo))

			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("Z", rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("I", rowNo))

					// URL
					w.Set(s.Cell("J", rowNo), s.URL(a.Url))
					w.CellExternalHyperLink(s.Cell("J", rowNo), s.URL(a.Url))
					w.CellStyle(s.Cell("J", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("J", rowNo), s.Cell("Z", rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
				s.Borders(s.FullBorder),
			))
			w.Merge(s.Cell("A", rowNo), s.Cell("I", rowNo))

			// URL
			w.Set(s.Cell("J", rowNo), s.GetConfigValue(cfg, s.CvWeb))
			w.CellExternalHyperLink(s.Cell("J", rowNo), s.GetConfigValue(cfg, s.CvWeb))
			w.CellStyle(s.Cell("J", rowNo), s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(s.Cell("J", rowNo), s.Cell("Z", rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * PRラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)
			w.Set(s.Cell("A", rowNo), "PR")
			w.Merge(s.Cell("A", rowNo), s.Cell("Z", rowNo))
			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("Z", rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
				s.Borders(s.FullBorder),
			))
			w.Merge(s.Cell("A", rowNo), s.Cell("Z", rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))

			w.InsertPageBreak(s.Cell("AB", rowNo+1))
		}

		/*
		 * スキルラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)
			w.Set(s.Cell("A", rowNo), "スキル")
			w.Merge(s.Cell("A", rowNo), s.Cell("Z", rowNo))
			w.HeaderCellRangeStyle(s.Cell("A", rowNo), s.Cell("Z", rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
					w.Height(rowNo, 140.0)

					w.Set(s.Cell("A", rowNo), *t.TagName)
					w.CellStyle(s.Cell("A", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(s.Cell("A", rowNo), s.Cell("G", rowNo))

					if t.Skills != nil {
						skillsBuf := strings.Builder{}
						for idx, sk := range *t.Skills {
							if idx != 0 {
								skillsBuf.WriteString(", ")
							}
							skillsBuf.WriteString(*sk.Name)

							if sk.Versions != nil {
								skillsBuf.WriteString("（経験：")
								var dms int
								for _, v := range *sk.Versions {
									if v.From == nil || v.To == nil {
										continue
									}
									dms += diffMonths(*v.From.Year, *v.From.Month, *v.To.Year, *v.To.Month)
								}
								skillsBuf.WriteString(mergeMonths(dms))
								skillsBuf.WriteString("）")
							}
						}
						w.Set(s.Cell("H", rowNo), skillsBuf.String())
						w.CellStyle(s.Cell("H", rowNo), s.NewStyle(
							s.Alignment(s.HLeftAlignment),
							s.Borders(s.FullBorder),
						))
					}
					w.Merge(s.Cell("H", rowNo), s.Cell("Z", rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell("AA", rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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

func diffMonths(fromYear, fromMonth, toYear, toMonth int) int {
	if fromYear == toYear {
		return toMonth - fromMonth + 1
	}

	diffMonth := 12 - fromMonth + 1 + toMonth

	diffYear := toYear - fromYear
	if diffYear == 1 {
		return diffMonth
	}

	return (diffYear-1)*12 + diffMonth
}

func mergeMonths(dms int) string {
	if dms < 12 {
		return fmt.Sprintf("%dヶ月", dms)
	}

	// FIXME:
	return ""
}
