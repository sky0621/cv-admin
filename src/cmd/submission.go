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

		careerGroups, err := requestUserInfo(cli, fmt.Sprintf("users/%d/careergroups", targetUserID), &[]rest.UserCareerGroup{})
		if err != nil {
			os.Exit(1)
		}

		/* **********************************************************
		 * Excelに書き込み
		 */
		w := s.NewExcelizeWrapper(
			s.SheetName(s.SkillSheetTitle),
			// MEMO: パスワードロックしても別シート作ってコピペはできるので（誤操作を防ぎたいわけではないので）意味なし。。
			// s.SheetPassword(s.GetConfigValue(cfg, s.Password)),
			s.PaperSize(9), // A4 210 x 297 mm
			s.SheetView(),
		)

		/*
		 * 列ごとの設定
		 */
		{
			w.Width(s.StartCol, s.EndCol, 5)
			w.Width(s.SuppleCol, s.SuppleCol, 0.5)
		}

		/*
		 * タイトル
		 */
		{
			w.Height(s.TitleRow, 40)

			titleCell := s.Cell(s.StartCol, s.TitleRow)
			w.Set(titleCell, s.SkillSheetTitle)
			w.CellStyle(titleCell, s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
				s.Font(s.SheetTitleFont),
			))
			w.Merge(titleCell, s.Cell(s.EndCol, s.TitleRow))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, s.TitleRow), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 最終更新日
		 */
		{
			w.Height(s.LastUpdatedRow, 20)

			lastUpdatedCell := s.Cell(s.StartCol, s.LastUpdatedRow)
			w.Set(lastUpdatedCell, fmt.Sprintf("最終更新日：%v", time.Now().Format("2006-01-02")))
			w.CellStyle(lastUpdatedCell, s.NewStyle(
				s.Alignment(s.HRightAlignment),
				s.Font(s.LastUpdatedFont),
			))
			w.Merge(lastUpdatedCell, s.Cell(s.EndCol, s.LastUpdatedRow))
		}

		/*
		 * 自動生成文言
		 */
		{
			w.Height(s.SkillSheetNoteRow, 20)

			skillSheetNoteCell := s.Cell(s.StartCol, s.SkillSheetNoteRow)
			w.Set(skillSheetNoteCell, "当スキルシートは以下の機能にて自動生成")
			w.CellStyle(skillSheetNoteCell, s.NewStyle(
				s.Alignment(s.HRightAlignment),
				s.Font(s.LastUpdatedFont),
			))
			w.Merge(skillSheetNoteCell, s.Cell(s.EndCol, s.SkillSheetNoteRow))
		}

		/*
		 * 自動生成プログラムURL
		 */
		{
			w.Height(s.SkillSheetNoteRow2, 20)

			skillSheetNoteCell2 := s.Cell(s.StartCol, s.SkillSheetNoteRow2)
			cvAdminURL := s.GetConfigValue(cfg, s.CvAdmin)
			w.Set(skillSheetNoteCell2, cvAdminURL)
			w.CellExternalHyperLink(skillSheetNoteCell2, cvAdminURL)
			w.CellStyle(skillSheetNoteCell2, s.NewStyle(
				s.Alignment(s.HRightAlignment),
				s.Font(s.LastUpdatedFont),
			))
			w.Merge(skillSheetNoteCell2, s.Cell(s.EndCol, s.SkillSheetNoteRow2))
		}

		/*
		 * 基礎属性１
		 */
		{
			w.Height(s.BasicAttributeRow, s.RowBaseHeight)

			kanaLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow)
			w.Set(kanaLabelCell, "フリガナ")
			w.HeaderCellStyle(kanaLabelCell)
			w.Merge(kanaLabelCell, s.Cell("C", s.BasicAttributeRow))

			kanaCell := s.Cell("D", s.BasicAttributeRow)
			w.Set(kanaCell, s.GetConfigValue(cfg, s.Kana))
			w.CellStyle(kanaCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(kanaCell, s.Cell("I", s.BasicAttributeRow))

			nicknameLabelCell := s.Cell("J", s.BasicAttributeRow)
			w.Set(nicknameLabelCell, "ニックネーム")
			w.Merge(nicknameLabelCell, s.Cell("M", s.BasicAttributeRow))

			ageLabelCell := s.Cell("N", s.BasicAttributeRow)
			w.Set(ageLabelCell, "年齢")
			w.Merge(ageLabelCell, s.Cell("O", s.BasicAttributeRow))

			mailLabelCell := s.Cell("P", s.BasicAttributeRow)
			w.Set(mailLabelCell, "メールアドレス")
			w.Merge(mailLabelCell, s.Cell("W", s.BasicAttributeRow))

			gravatarLabelCell := s.Cell("X", s.BasicAttributeRow)
			w.Set(gravatarLabelCell, "Gravatar")
			w.Merge(gravatarLabelCell, s.Cell(s.EndCol, s.BasicAttributeRow))

			w.HeaderCellRangeStyle(nicknameLabelCell, s.Cell(s.EndCol, s.BasicAttributeRow))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, s.BasicAttributeRow), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 基礎属性２
		 */
		{
			w.Height(s.BasicAttributeRow2, s.RowBaseHeight*2)

			nameLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow2)
			w.Set(nameLabelCell, "名前")
			w.HeaderCellStyle(nameLabelCell)
			w.Merge(nameLabelCell, s.Cell("C", s.BasicAttributeRow2))

			nameCell := s.Cell("D", s.BasicAttributeRow2)
			w.Set(nameCell, s.GetConfigValue(cfg, s.Name))
			w.CellStyle(nameCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
				s.Font(s.NameFont),
			))
			w.Merge(nameCell, s.Cell("I", s.BasicAttributeRow2))

			nicknameCell := s.Cell("J", s.BasicAttributeRow2)
			w.Set(nicknameCell, *attribute.Nickname)
			w.CellStyle(nicknameCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(nicknameCell, s.Cell("M", s.BasicAttributeRow2))

			ageCell := s.Cell("N", s.BasicAttributeRow2)
			bDay := *attribute.Birthday
			w.Set(ageCell, s.Age(*bDay.Year, *bDay.Month, *bDay.Day, time.Now()))
			w.CellStyle(ageCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(ageCell, s.Cell("O", s.BasicAttributeRow2))

			mailCell := s.Cell("P", s.BasicAttributeRow2)
			w.Set(mailCell, s.GetConfigValue(cfg, s.Mail))
			w.CellStyle(mailCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(mailCell, s.Cell("W", s.BasicAttributeRow2))

			gravatarCell := s.Cell("X", s.BasicAttributeRow2)
			if avatarImgBytes != nil {
				w.AddPictureFromBytes(gravatarCell, "avatar", ".png", avatarImgBytes)
			}
			w.CellStyle(gravatarCell, s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
			))

			w.Merge(gravatarCell, s.Cell(s.EndCol, s.BasicAttributeRow2))
			w.Merge(s.Cell("X", s.BasicAttributeRow3), s.Cell(s.EndCol, s.BasicAttributeRow3))
			w.Merge(s.Cell("X", s.BasicAttributeRow4), s.Cell(s.EndCol, s.BasicAttributeRow4))

			w.Merge(gravatarCell, s.Cell("X", s.BasicAttributeRow4))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, s.BasicAttributeRow2), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 基礎属性３
		 */
		{
			w.Height(s.BasicAttributeRow3, s.RowBaseHeight)

			residenceLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow3)
			w.Set(residenceLabelCell, "居住地")
			w.HeaderCellStyle(residenceLabelCell)
			w.Merge(residenceLabelCell, s.Cell("C", s.BasicAttributeRow3))

			residenceCell := s.Cell("D", s.BasicAttributeRow3)
			w.Set(residenceCell, s.GetConfigValue(cfg, s.CityOfResidence))
			w.CellStyle(residenceCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(residenceCell, s.Cell("I", s.BasicAttributeRow3))

			nearStLabelCell := s.Cell("J", s.BasicAttributeRow3)
			w.Set(nearStLabelCell, "最寄り駅")
			w.HeaderCellStyle(nearStLabelCell)
			w.Merge(nearStLabelCell, s.Cell("M", s.BasicAttributeRow3))

			nearStCell := s.Cell("N", s.BasicAttributeRow3)
			w.Set(nearStCell, s.GetConfigValue(cfg, s.NearestStation))
			w.CellStyle(nearStCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(nearStCell, s.Cell("W", s.BasicAttributeRow3))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, s.BasicAttributeRow3), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 基礎属性４
		 */
		{
			w.Height(s.BasicAttributeRow4, s.RowBaseHeight)

			eduBgLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow4)
			w.Set(eduBgLabelCell, "最終学歴")
			w.HeaderCellStyle(eduBgLabelCell)
			w.Merge(eduBgLabelCell, s.Cell("C", s.BasicAttributeRow4))

			eduBgCell := s.Cell("D", s.BasicAttributeRow4)
			w.Set(eduBgCell, s.GetConfigValue(cfg, s.EducationalBackground))
			w.CellStyle(eduBgCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(eduBgCell, s.Cell("W", s.BasicAttributeRow4))

			// 枠線「右」「下」が機能していないための措置
			w.CellStyle(s.Cell("X", s.BasicAttributeRow4), s.NewStyle(s.Borders(s.BottomBorder)))
			w.CellStyle(s.Cell(s.SuppleCol, s.BasicAttributeRow4), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 資格情報ラベル
		 */
		{
			w.Height(s.QualificationRow, s.RowBaseHeight)

			qualificationLabelCell := s.Cell(s.StartCol, s.QualificationRow)
			w.Set(qualificationLabelCell, "資格")
			w.Merge(qualificationLabelCell, s.Cell("I", s.QualificationRow))

			qualificationGotDateLabelCell := s.Cell("J", s.QualificationRow)
			w.Set(qualificationGotDateLabelCell, "取得年月日")
			w.Merge(qualificationGotDateLabelCell, s.Cell("M", s.QualificationRow))

			qualificationURLLabelCell := s.Cell("N", s.QualificationRow)
			w.Set(qualificationURLLabelCell, "URL")
			w.Merge(qualificationURLLabelCell, s.Cell(s.EndCol, s.QualificationRow))

			w.HeaderCellRangeStyle(qualificationLabelCell, s.Cell(s.EndCol, s.QualificationRow))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, s.QualificationRow), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * 資格情報
		 */
		rowNo := s.QualificationRow
		{
			if qualifications != nil {
				for _, q := range *qualifications {
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight)

					// 資格
					qualificationNameCell := s.Cell(s.StartCol, rowNo)
					w.Set(qualificationNameCell, s.QualificationName(q.Organization, q.Name))
					w.CellStyle(qualificationNameCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(qualificationNameCell, s.Cell("I", rowNo))

					// 取得年月日
					qualificationGotDateCell := s.Cell("J", rowNo)
					w.Set(qualificationGotDateCell, s.QualificationGotDate(q.GotDate))
					w.CellStyle(qualificationGotDateCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(qualificationGotDateCell, s.Cell("M", rowNo))

					// URL
					qualificationURLCell := s.Cell("N", rowNo)
					w.Set(qualificationURLCell, s.URL(q.Url))
					w.CellExternalHyperLink(qualificationURLCell, s.URL(q.Url))
					w.CellStyle(qualificationURLCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(qualificationURLCell, s.Cell(s.EndCol, rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
				}
			}
		}

		/*
		 * アウトプットラベル
		 */
		rowNo += 2
		{
			w.Height(rowNo, s.RowBaseHeight)

			outputLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(outputLabelCell, "アウトプット")
			w.Merge(outputLabelCell, s.Cell("I", rowNo))

			outputURLCell := s.Cell("J", rowNo)
			w.Set(outputURLCell, "URL")
			w.Merge(outputURLCell, s.Cell("Z", rowNo))

			w.HeaderCellRangeStyle(outputLabelCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
					outputNameCell := s.Cell(s.StartCol, rowNo)
					w.Set(outputNameCell, *a.Name)
					w.CellStyle(outputNameCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(outputNameCell, s.Cell("I", rowNo))

					// URL
					outputURLCell := s.Cell("J", rowNo)
					w.Set(outputURLCell, s.URL(a.Url))
					w.CellExternalHyperLink(outputURLCell, s.URL(a.Url))
					w.CellStyle(outputURLCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(outputURLCell, s.Cell(s.EndCol, rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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
			outputNameCell := s.Cell(s.StartCol, rowNo)
			w.Set(outputNameCell, "スキルシート(Web版)")
			w.CellStyle(outputNameCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(outputNameCell, s.Cell("I", rowNo))

			// URL
			outputURLCell := s.Cell("J", rowNo)
			w.Set(outputURLCell, s.GetConfigValue(cfg, s.CvWeb))
			w.CellExternalHyperLink(outputURLCell, s.GetConfigValue(cfg, s.CvWeb))
			w.CellStyle(outputURLCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(outputURLCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * PRラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)

			prLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(prLabelCell, "PR")
			w.Merge(prLabelCell, s.Cell(s.EndCol, rowNo))
			w.HeaderCellRangeStyle(prLabelCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * PR
		 */
		{
			rowNo += 1
			w.Height(rowNo, s.RowMaxHeight) // TODO: PR行数に応じて（高さが収まらない場合があるので）セルを分ける対応が必要

			prCell := s.Cell(s.StartCol, rowNo)
			w.Set(prCell, *attribute.Pr)
			w.CellStyle(prCell, s.NewStyle(
				s.Alignment(s.HLeftAlignment),
				s.Borders(s.FullBorder),
			))
			w.Merge(prCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))

			w.InsertPageBreak(s.Cell("AB", rowNo+1))
		}

		/*
		 * スキルラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)

			skillLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(skillLabelCell, "スキル")
			w.Merge(skillLabelCell, s.Cell(s.EndCol, rowNo))
			w.HeaderCellRangeStyle(skillLabelCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
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

					tagNameCell := s.Cell(s.StartCol, rowNo)
					w.Set(tagNameCell, *t.TagName)
					w.CellStyle(tagNameCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.FullBorder),
					))
					w.Merge(tagNameCell, s.Cell("G", rowNo))

					skillCell := ""
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
									dms += s.DiffMonths(*v.From.Year, *v.From.Month, *v.To.Year, *v.To.Month)
								}
								skillsBuf.WriteString(s.MergeMonths(dms))
								skillsBuf.WriteString("）")
							}
						}
						skillCell = s.Cell("H", rowNo)
						w.Set(skillCell, skillsBuf.String())
						w.CellStyle(skillCell, s.NewStyle(
							s.Alignment(s.HLeftAlignment),
							s.Borders(s.FullBorder),
						))
					}
					w.Merge(skillCell, s.Cell(s.EndCol, rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
				}
			}
		}

		/*
		 * キャリアラベル
		 */
		{
			rowNo += 2
			w.Height(rowNo, s.RowBaseHeight)

			careerLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(careerLabelCell, "キャリア")
			w.Merge(careerLabelCell, s.Cell(s.EndCol, rowNo))
			w.HeaderCellRangeStyle(careerLabelCell, s.Cell(s.EndCol, rowNo))

			// 枠線「右」が機能していないための措置
			w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
		}

		/*
		 * キャリア
		 */
		{
			if careerGroups != nil {
				for idx, cg := range *careerGroups {
					/*
					 * キャリアグループラベル
					 */
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight*1.8)

					careerGroupLabelCell := s.Cell(s.StartCol, rowNo)
					w.Set(careerGroupLabelCell, fmt.Sprintf("(%d) %s", idx+1, *cg.Label))
					w.CellStyle(careerGroupLabelCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Fill(s.CareerGroupLabelFill),
						s.Font(s.CareerGroupLabelFont),
					))
					w.Merge(careerGroupLabelCell, s.Cell("P", rowNo))

					careerGroupPeriodLabelCell := s.Cell("Q", rowNo)
					w.Set(careerGroupPeriodLabelCell, "期間：")
					w.CellStyle(careerGroupPeriodLabelCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Fill(s.CareerGroupLabelFill),
					))
					w.Merge(careerGroupPeriodLabelCell, s.Cell("R", rowNo))

					careerGroupPeriodCell := s.Cell("S", rowNo)
					w.Set(careerGroupPeriodCell, s.CareerGroupPeriod(cg))
					w.CellStyle(careerGroupPeriodCell, s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Fill(s.CareerGroupLabelFill),
					))
					w.Merge(careerGroupPeriodCell, s.Cell(s.EndCol, rowNo))

					// 枠線「右」が機能していないための措置
					w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))

					/*
					 * キャリア群
					 */
					if cg.Careers != nil {
						for idx2, c := range *cg.Careers {
							/*
							 * キャリアタイトル
							 */
							rowNo += 1
							w.Height(rowNo, s.RowBaseHeight*1.5)

							careerNameCell := s.Cell(s.StartCol, rowNo)
							w.Set(careerNameCell, fmt.Sprintf("(%d - %d) %s", idx+1, idx2+1, *c.Name))
							w.CellStyle(careerNameCell, s.NewStyle(
								s.Alignment(s.HLeftAlignmentIndent2),
								s.Borders(s.FullBorder),
								s.Font(s.CareerNameFont),
							))
							w.Merge(careerNameCell, s.Cell("P", rowNo))

							careerPeriodLabelCell := s.Cell("Q", rowNo)
							w.Set(careerPeriodLabelCell, "期間：")
							w.HeaderCellRangeStyle(careerPeriodLabelCell, s.Cell("R", rowNo))
							w.Merge(careerPeriodLabelCell, s.Cell("R", rowNo))

							careerPeriodCell := s.Cell("S", rowNo)
							w.Set(careerPeriodCell, s.CareerPeriod(c))
							w.CellStyle(careerPeriodCell, s.NewStyle(
								s.Alignment(s.HLeftAlignmentIndent2),
								s.Borders(s.FullBorder),
							))
							w.Merge(careerPeriodCell, s.Cell(s.EndCol, rowNo))

							// 枠線「右」が機能していないための措置
							w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))

							/*
							 * 説明
							 */
							if c.Description != nil {
								for _, desc := range *c.Description {
									if desc != "-" {
										rowNo += 1
										w.Height(rowNo, s.RowBaseHeight*1.5)

										careerDescCell := s.Cell(s.StartCol, rowNo)
										w.Set(careerDescCell, desc)
										w.CellStyle(careerDescCell, s.NewStyle(
											s.Alignment(s.HLeftAlignmentIndent3),
											s.Borders(s.SideBorder),
										))
										w.Merge(careerDescCell, s.Cell(s.EndCol, rowNo))

										// 枠線「右」が機能していないための措置
										w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
									}
								}
							}

							/*
							 * タスク
							 */
							if c.Tasks != nil {
								rowNo += 1
								w.Height(rowNo, s.RowBaseHeight)

								taskLabelCell := s.Cell(s.StartCol, rowNo)
								w.Set(taskLabelCell, "タスク")
								w.CellStyle(taskLabelCell, s.NewStyle(
									s.Alignment(s.HLeftAlignmentIndent3),
									s.Borders(s.SideBorder),
									s.Font(s.BoldFont),
								))
								w.Merge(taskLabelCell, s.Cell(s.EndCol, rowNo))

								// 枠線「右」が機能していないための措置
								w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))

								for _, task := range *c.Tasks {
									if task.Name != nil && *task.Name != "-" {
										rowNo += 1
										w.Height(rowNo, s.RowBaseHeight)

										taskNameCell := s.Cell(s.StartCol, rowNo)
										w.Set(taskNameCell, fmt.Sprintf("【%s】", *task.Name))
										w.CellStyle(taskNameCell, s.NewStyle(
											s.Alignment(s.HLeftAlignmentIndent3),
											s.Borders(s.SideBorder),
										))
										w.Merge(taskNameCell, s.Cell(s.EndCol, rowNo))

										// 枠線「右」が機能していないための措置
										w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
									}

									if task.Description != nil {
										for _, desc := range *task.Description {
											rowNo += 1
											w.Height(rowNo, s.RowBaseHeight)

											taskNameCell := s.Cell(s.StartCol, rowNo)
											w.Set(taskNameCell, desc)
											w.CellStyle(taskNameCell, s.NewStyle(
												s.Alignment(s.HLeftAlignmentIndent4),
												s.Borders(s.SideBorder),
											))
											w.Merge(taskNameCell, s.Cell(s.EndCol, rowNo))

											// 枠線「右」が機能していないための措置
											w.CellStyle(s.Cell(s.SuppleCol, rowNo), s.NewStyle(s.Borders(s.LeftBorder)))
										}
									}
								}
							}
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
