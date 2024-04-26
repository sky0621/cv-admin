/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/sky0621/cv-admin/src/rest"
	s "github.com/sky0621/cv-admin/src/submission"
	util "github.com/sky0621/golang-utils/string"
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
		if attribute != nil && attribute.AvatarUrl != nil {
			avatarImgBytes, err = requestAvatarImage(cli, *attribute.AvatarUrl)
			if err != nil {
				os.Exit(1)
			}
		}

		solutions, err := requestUserInfo(cli, fmt.Sprintf("users/%d/solutions", targetUserID), &[]rest.UserSolution{})
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(solutions)

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
			w.Width(s.StartCol, s.EndCol, s.BaseColWidth)
			w.Width(s.SuppleCol, s.SuppleCol, s.SuppleColWidth)
		}

		/*
		 * タイトル
		 */
		{
			w.Height(s.TitleRow, 40)

			titleCell := s.Cell(s.StartCol, s.TitleRow)
			w.Set(titleCell, s.SkillSheetTitle)
			w.Merge(titleCell, s.Cell(s.EndCol, s.TitleRow))
			w.CellRangeStyle(titleCell, s.Cell(s.EndCol, s.TitleRow), s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
				s.Font(s.SheetTitleFont),
			))
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

			/*
			 * 「フリガナ」ラベル
			 */
			kanaLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow)
			w.Set(kanaLabelCell, "フリガナ")
			w.Merge(kanaLabelCell, s.Cell("C", s.BasicAttributeRow))
			w.HeaderCellRangeStyle(kanaLabelCell, s.Cell("C", s.BasicAttributeRow))

			/*
			 * 「フリガナ」
			 */
			kanaCell := s.Cell("D", s.BasicAttributeRow)
			w.Set(kanaCell, s.GetConfigValue(cfg, s.Kana))
			w.Merge(kanaCell, s.Cell("I", s.BasicAttributeRow))
			w.ValueCellRangeStyle(kanaCell, s.Cell("I", s.BasicAttributeRow))

			/*
			 * 「ニックネーム」ラベル
			 */
			nicknameLabelCell := s.Cell("J", s.BasicAttributeRow)
			w.Set(nicknameLabelCell, "ニックネーム")
			w.Merge(nicknameLabelCell, s.Cell("M", s.BasicAttributeRow))
			w.HeaderCellRangeStyle(nicknameLabelCell, s.Cell("M", s.BasicAttributeRow))

			/*
			 * 「年齢」ラベル
			 */
			ageLabelCell := s.Cell("N", s.BasicAttributeRow)
			w.Set(ageLabelCell, "年齢")
			w.Merge(ageLabelCell, s.Cell("O", s.BasicAttributeRow))
			w.HeaderCellRangeStyle(ageLabelCell, s.Cell("O", s.BasicAttributeRow))

			/*
			 * 「メールアドレス」ラベル
			 */
			mailLabelCell := s.Cell("P", s.BasicAttributeRow)
			w.Set(mailLabelCell, "メールアドレス")
			w.Merge(mailLabelCell, s.Cell("W", s.BasicAttributeRow))
			w.HeaderCellRangeStyle(mailLabelCell, s.Cell("W", s.BasicAttributeRow))

			/*
			 * 「Gravatar」ラベル
			 */
			gravatarLabelCell := s.Cell("X", s.BasicAttributeRow)
			w.Set(gravatarLabelCell, "Gravatar")
			w.Merge(gravatarLabelCell, s.Cell(s.EndCol, s.BasicAttributeRow))
			w.HeaderCellRangeStyle(gravatarLabelCell, s.Cell(s.EndCol, s.BasicAttributeRow))
		}

		/*
		 * 基礎属性２
		 */
		{
			w.Height(s.BasicAttributeRow2, s.RowBaseHeight*2)

			/*
			 *「名前」ラベル
			 */
			nameLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow2)
			w.Set(nameLabelCell, "名前")
			w.Merge(nameLabelCell, s.Cell("C", s.BasicAttributeRow2))
			w.HeaderCellRangeStyle(nameLabelCell, s.Cell("C", s.BasicAttributeRow2))

			/*
			 * 「名前」
			 */
			nameCell := s.Cell("D", s.BasicAttributeRow2)
			w.Set(nameCell, s.GetConfigValue(cfg, s.Name))
			w.Merge(nameCell, s.Cell("I", s.BasicAttributeRow2))
			w.ValueCellRangeStyle(nameCell, s.Cell("I", s.BasicAttributeRow2))

			/*
			 * 「ニックネーム」
			 */
			nicknameCell := s.Cell("J", s.BasicAttributeRow2)
			w.Set(nicknameCell, *attribute.Nickname)
			w.Merge(nicknameCell, s.Cell("M", s.BasicAttributeRow2))
			w.ValueCellRangeStyle(nicknameCell, s.Cell("M", s.BasicAttributeRow2))

			/*
			 * 「年齢」
			 */
			ageCell := s.Cell("N", s.BasicAttributeRow2)
			bDay := *attribute.Birthday
			w.Set(ageCell, s.Age(*bDay.Year, *bDay.Month, *bDay.Day, time.Now()))
			w.Merge(ageCell, s.Cell("O", s.BasicAttributeRow2))
			w.ValueCellRangeStyle(ageCell, s.Cell("O", s.BasicAttributeRow2))

			/*
			 * 「メールアドレス」
			 */
			mailCell := s.Cell("P", s.BasicAttributeRow2)
			w.Set(mailCell, s.GetConfigValue(cfg, s.Mail))
			w.Merge(mailCell, s.Cell("W", s.BasicAttributeRow2))
			w.ValueCellRangeStyle(mailCell, s.Cell("W", s.BasicAttributeRow2))

			/*
			 * 「Gravatar」
			 */
			gravatarCell := s.Cell("X", s.BasicAttributeRow2)
			if avatarImgBytes != nil {
				w.AddPictureFromBytes(gravatarCell, ".png", avatarImgBytes)
			}

			w.Merge(gravatarCell, s.Cell(s.EndCol, s.BasicAttributeRow2))
			w.Merge(s.Cell("X", s.BasicAttributeRow3), s.Cell(s.EndCol, s.BasicAttributeRow3))
			w.Merge(s.Cell("X", s.BasicAttributeRow4), s.Cell(s.EndCol, s.BasicAttributeRow4))

			w.Merge(gravatarCell, s.Cell("X", s.BasicAttributeRow4))
			w.CellRangeStyle(gravatarCell, s.Cell(s.EndCol, s.BasicAttributeRow4), s.NewStyle(
				s.Alignment(s.VhCenterAlignment),
				s.Borders(s.FullBorder),
			))
		}

		/*
		 * 基礎属性３
		 */
		{
			w.Height(s.BasicAttributeRow3, s.RowBaseHeight)

			/*
			 * 「居住地」ラベル
			 */
			residenceLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow3)
			w.Set(residenceLabelCell, "居住地")
			w.Merge(residenceLabelCell, s.Cell("C", s.BasicAttributeRow3))
			w.HeaderCellRangeStyle(residenceLabelCell, s.Cell("C", s.BasicAttributeRow3))

			/*
			 * 「居住地」
			 */
			residenceCell := s.Cell("D", s.BasicAttributeRow3)
			w.Set(residenceCell, s.GetConfigValue(cfg, s.CityOfResidence))
			w.Merge(residenceCell, s.Cell("I", s.BasicAttributeRow3))
			w.ValueCellRangeStyle(residenceCell, s.Cell("I", s.BasicAttributeRow3))

			/*
			 * 「最寄り駅」ラベル
			 */
			nearStLabelCell := s.Cell("J", s.BasicAttributeRow3)
			w.Set(nearStLabelCell, "最寄り駅")
			w.Merge(nearStLabelCell, s.Cell("M", s.BasicAttributeRow3))
			w.HeaderCellRangeStyle(nearStLabelCell, s.Cell("M", s.BasicAttributeRow3))

			/*
			 * 「最寄り駅」
			 */
			nearStCell := s.Cell("N", s.BasicAttributeRow3)
			w.Set(nearStCell, s.GetConfigValue(cfg, s.NearestStation))
			w.Merge(nearStCell, s.Cell("W", s.BasicAttributeRow3))
			w.ValueCellRangeStyle(nearStCell, s.Cell("W", s.BasicAttributeRow3))
		}

		/*
		 * 基礎属性４
		 */
		{
			w.Height(s.BasicAttributeRow4, s.RowBaseHeight)

			/*
			 * 「最終学歴」ラベル
			 */
			eduBgLabelCell := s.Cell(s.StartCol, s.BasicAttributeRow4)
			w.Set(eduBgLabelCell, "最終学歴")
			w.Merge(eduBgLabelCell, s.Cell("C", s.BasicAttributeRow4))
			w.HeaderCellRangeStyle(eduBgLabelCell, s.Cell("C", s.BasicAttributeRow4))

			/*
			 * 「最終学歴」
			 */
			eduBgCell := s.Cell("D", s.BasicAttributeRow4)
			w.Set(eduBgCell, s.GetConfigValue(cfg, s.EducationalBackground))
			w.Merge(eduBgCell, s.Cell("W", s.BasicAttributeRow4))
			w.ValueCellRangeStyle(eduBgCell, s.Cell("W", s.BasicAttributeRow4))
		}

		/*
		 * 資格情報ラベル
		 */
		{
			w.Height(s.QualificationRow, s.RowBaseHeight)

			/*
			 * 「資格」ラベル
			 */
			qualificationLabelCell := s.Cell(s.StartCol, s.QualificationRow)
			w.Set(qualificationLabelCell, "資格")
			w.Merge(qualificationLabelCell, s.Cell("I", s.QualificationRow))
			w.HeaderCellRangeStyle(qualificationLabelCell, s.Cell("I", s.QualificationRow))

			/*
			 * 「取得年月日」ラベル
			 */
			qualificationGotDateLabelCell := s.Cell("J", s.QualificationRow)
			w.Set(qualificationGotDateLabelCell, "取得年月日")
			w.Merge(qualificationGotDateLabelCell, s.Cell("M", s.QualificationRow))
			w.HeaderCellRangeStyle(qualificationGotDateLabelCell, s.Cell("M", s.QualificationRow))

			/*
			 * 「URL」ラベル
			 */
			qualificationURLLabelCell := s.Cell("N", s.QualificationRow)
			w.Set(qualificationURLLabelCell, "URL")
			w.Merge(qualificationURLLabelCell, s.Cell(s.EndCol, s.QualificationRow))
			w.HeaderCellRangeStyle(qualificationURLLabelCell, s.Cell(s.EndCol, s.QualificationRow))
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

					/*
					 * 「資格」
					 */
					qualificationNameCell := s.Cell(s.StartCol, rowNo)
					w.Set(qualificationNameCell, s.QualificationName(q.Organization, q.Name))
					w.Merge(qualificationNameCell, s.Cell("I", rowNo))
					w.ValueCellRangeStyle(qualificationNameCell, s.Cell("I", rowNo))

					/*
					 * 「取得年月日」
					 */
					qualificationGotDateCell := s.Cell("J", rowNo)
					w.Set(qualificationGotDateCell, s.QualificationGotDate(q.GotDate))
					w.Merge(qualificationGotDateCell, s.Cell("M", rowNo))
					w.ValueCellRangeStyle(qualificationGotDateCell, s.Cell("M", rowNo))

					/*
					 * 「URL」
					 */
					qualificationURLCell := s.Cell("N", rowNo)
					w.Set(qualificationURLCell, s.URL(q.Url))
					w.CellExternalHyperLink(qualificationURLCell, s.URL(q.Url))
					w.Merge(qualificationURLCell, s.Cell(s.EndCol, rowNo))
					w.ValueCellRangeStyle(qualificationURLCell, s.Cell(s.EndCol, rowNo))
				}
			}
		}

		/*
		 * アウトプットラベル
		 */
		rowNo += 2
		{
			w.Height(rowNo, s.RowBaseHeight)

			/*
			 * 「アウトプット」ラベル
			 */
			outputLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(outputLabelCell, "アウトプット")
			w.Merge(outputLabelCell, s.Cell("I", rowNo))
			w.HeaderCellRangeStyle(outputLabelCell, s.Cell("I", rowNo))

			/*
			 * 「URL」ラベル
			 */
			outputURLCell := s.Cell("J", rowNo)
			w.Set(outputURLCell, "URL")
			w.Merge(outputURLCell, s.Cell(s.EndCol, rowNo))
			w.HeaderCellRangeStyle(outputURLCell, s.Cell(s.EndCol, rowNo))
		}

		/*
		 * アウトプット
		 */
		{
			if activities != nil {
				for _, a := range *activities {
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight)

					/*
					 * 「アウトプット」
					 */
					outputNameCell := s.Cell(s.StartCol, rowNo)
					w.Set(outputNameCell, *a.Name)
					w.Merge(outputNameCell, s.Cell("I", rowNo))
					w.ValueCellRangeStyle(outputNameCell, s.Cell("I", rowNo))

					/*
					 * 「URL」
					 */
					outputURLCell := s.Cell("J", rowNo)
					w.Set(outputURLCell, s.URL(a.Url))
					w.CellExternalHyperLink(outputURLCell, s.URL(a.Url))
					w.Merge(outputURLCell, s.Cell(s.EndCol, rowNo))
					w.ValueCellRangeStyle(outputURLCell, s.Cell(s.EndCol, rowNo))
				}
			}
		}

		/*
		 * アウトプット（from .private.ini）
		 */
		{
			rowNo += 1
			w.Height(rowNo, s.RowBaseHeight)

			/*
			 * 「アウトプット」
			 */
			outputNameCell := s.Cell(s.StartCol, rowNo)
			w.Set(outputNameCell, "スキルシート(Web版)")
			w.Merge(outputNameCell, s.Cell("I", rowNo))
			w.ValueCellRangeStyle(outputNameCell, s.Cell("I", rowNo))

			/*
			 * 「URL」
			 */
			outputURLCell := s.Cell("J", rowNo)
			w.Set(outputURLCell, s.GetConfigValue(cfg, s.CvWeb))
			w.CellExternalHyperLink(outputURLCell, s.GetConfigValue(cfg, s.CvWeb))
			w.Merge(outputURLCell, s.Cell(s.EndCol, rowNo))
			w.ValueCellRangeStyle(outputURLCell, s.Cell(s.EndCol, rowNo))
		}

		/*
		 * PRラベル
		 */
		rowNo += 2
		{
			w.Height(rowNo, s.RowBaseHeight)

			prLabelCell := s.Cell(s.StartCol, rowNo)
			w.Set(prLabelCell, "PR")
			w.Merge(prLabelCell, s.Cell(s.EndCol, rowNo))
			w.HeaderCellRangeStyle(prLabelCell, s.Cell(s.EndCol, rowNo))
		}

		/*
		 * PR
		 */
		{
			rowNo += 1
			w.Height(rowNo, s.RowMaxHeight) // TODO: PR行数に応じて（高さが収まらない場合があるので）セルを分ける対応が必要

			prCell := s.Cell(s.StartCol, rowNo)
			w.Set(prCell, *attribute.Pr)
			w.Merge(prCell, s.Cell(s.EndCol, rowNo))
			w.ValueCellRangeStyle(prCell, s.Cell(s.EndCol, rowNo))

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
					w.Merge(tagNameCell, s.Cell("G", rowNo))
					w.ValueCellRangeStyle(tagNameCell, s.Cell("G", rowNo))

					skillCell := ""
					if t.Skills != nil {
						builder := util.NewBuilder()
						for idx, sk := range *t.Skills {
							if idx != 0 {
								builder.Append(", ")
							}
							builder.Append(*sk.Name)

							if sk.Versions != nil {
								builder.Append("（経験：")
								var dms int
								for _, v := range *sk.Versions {
									if v.From == nil || v.To == nil {
										continue
									}
									dms += s.DiffMonths(*v.From.Year, *v.From.Month, *v.To.Year, *v.To.Month)
								}
								builder.Append(s.MergeMonths(dms))
								builder.Append("）")
							}
						}
						skillCell = s.Cell("H", rowNo)
						w.Set(skillCell, builder.ToString())
						w.CellStyle(skillCell, s.NewStyle(
							s.Alignment(s.HLeftAlignment),
							s.Borders(s.FullBorder),
						))
					}
					w.Merge(skillCell, s.Cell(s.EndCol, rowNo))
					w.ValueCellRangeStyle(skillCell, s.Cell(s.EndCol, rowNo))
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
		}

		/*
		 * キャリア
		 */
		{
			if careerGroups != nil {
				for idx, cg := range *careerGroups {
					rowNo += 1
					w.Height(rowNo, s.RowBaseHeight*1.8)

					/*
					 * 「キャリアグループタイトル」
					 */
					careerGroupLabelCell := s.Cell(s.StartCol, rowNo)
					w.Set(careerGroupLabelCell, fmt.Sprintf("(%d) %s", idx+1, *cg.Label))
					w.Merge(careerGroupLabelCell, s.Cell("P", rowNo))
					w.CellRangeStyle(careerGroupLabelCell, s.Cell("P", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.LeftTopBottomBorder),
						s.Fill(s.CareerGroupLabelFill),
						s.Font(s.CareerGroupLabelFont),
					))

					/*
					 * 「キャリアグループ期間」ラベル
					 */
					careerGroupPeriodLabelCell := s.Cell("Q", rowNo)
					w.Set(careerGroupPeriodLabelCell, "期間：")
					w.Merge(careerGroupPeriodLabelCell, s.Cell("R", rowNo))
					w.CellRangeStyle(careerGroupPeriodLabelCell, s.Cell("R", rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.TopBottomBorder),
						s.Fill(s.CareerGroupLabelFill),
					))

					/*
					 * 「キャリアグループ期間」
					 */
					careerGroupPeriodCell := s.Cell("S", rowNo)
					w.Set(careerGroupPeriodCell, s.CareerGroupPeriod(cg))
					w.Merge(careerGroupPeriodCell, s.Cell(s.EndCol, rowNo))
					w.CellRangeStyle(careerGroupPeriodCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
						s.Alignment(s.HLeftAlignment),
						s.Borders(s.RightTopBottomBorder),
						s.Fill(s.CareerGroupLabelFill),
					))

					/*
					 * キャリア群
					 */
					if cg.Careers != nil {
						for idx2, c := range *cg.Careers {
							rowNo += 1
							w.Height(rowNo, s.RowBaseHeight*1.5)

							/*
							 * 「キャリアタイトル」
							 */
							careerNameCell := s.Cell(s.StartCol, rowNo)
							w.Set(careerNameCell, fmt.Sprintf("(%d - %d) %s", idx+1, idx2+1, *c.Name))
							w.Merge(careerNameCell, s.Cell("P", rowNo))
							w.CellRangeStyle(careerNameCell, s.Cell("P", rowNo), s.NewStyle(
								s.Alignment(s.HLeftAlignmentIndent2),
								s.Borders(s.LeftTopBottomBorder),
								s.Font(s.CareerNameFont),
							))

							/*
							 * 「キャリア期間」ラベル
							 */
							careerPeriodLabelCell := s.Cell("Q", rowNo)
							w.Set(careerPeriodLabelCell, "期間：")
							w.Merge(careerPeriodLabelCell, s.Cell("R", rowNo))
							w.CellRangeStyle(careerPeriodLabelCell, s.Cell("R", rowNo), s.NewStyle(
								s.Alignment(s.HLeftAlignment),
								s.Borders(s.TopBottomBorder),
							))

							/*
							 * 「キャリア期間」
							 */
							careerPeriodCell := s.Cell("S", rowNo)
							w.Set(careerPeriodCell, s.CareerPeriod(c))
							w.Merge(careerPeriodCell, s.Cell(s.EndCol, rowNo))
							w.CellRangeStyle(careerPeriodCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
								s.Alignment(s.HLeftAlignment),
								s.Borders(s.RightTopBottomBorder),
							))

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
										w.Merge(careerDescCell, s.Cell(s.EndCol, rowNo))
										w.CellRangeStyle(careerDescCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
											s.Alignment(s.HLeftAlignmentIndent3),
											s.Borders(s.SideBorder),
										))
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
								w.Set(taskLabelCell, "担当タスク・役割")
								w.Merge(taskLabelCell, s.Cell(s.EndCol, rowNo))
								w.CellRangeStyle(taskLabelCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
									s.Alignment(s.HLeftAlignmentIndent3),
									s.Borders(s.SideBorder),
									s.Font(s.BoldFont),
								))

								for _, task := range *c.Tasks {
									if task.Name != nil && *task.Name != "-" {
										rowNo += 1
										w.Height(rowNo, s.RowBaseHeight)

										taskNameCell := s.Cell(s.StartCol, rowNo)
										w.Set(taskNameCell, fmt.Sprintf("【%s】", *task.Name))
										w.Merge(taskNameCell, s.Cell(s.EndCol, rowNo))
										w.CellRangeStyle(taskNameCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
											s.Alignment(s.HLeftAlignmentIndent3),
											s.Borders(s.SideBorder),
										))
									}

									if task.Description != nil {
										for _, desc := range *task.Description {
											rowNo += 1
											w.Height(rowNo, s.RowBaseHeight)

											taskNameCell := s.Cell(s.StartCol, rowNo)
											w.Set(taskNameCell, desc)
											w.Merge(taskNameCell, s.Cell(s.EndCol, rowNo))
											w.CellRangeStyle(taskNameCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
												s.Alignment(s.HLeftAlignmentIndent4),
												s.Borders(s.SideBorder),
											))
										}
									}
								}
							}

							/*
							 * スキル
							 */
							if c.SkillGroups != nil {
								rowNo += 1
								w.Height(rowNo, s.RowBaseHeight)

								skillLabelCell := s.Cell(s.StartCol, rowNo)
								w.Set(skillLabelCell, "使用技術")
								w.Merge(skillLabelCell, s.Cell(s.EndCol, rowNo))
								w.CellRangeStyle(skillLabelCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
									s.Alignment(s.HLeftAlignmentIndent3),
									s.Borders(s.SideBorder),
									s.Font(s.BoldFont),
								))

								for _, skillGroup := range *c.SkillGroups {
									if skillGroup.Label != nil && *skillGroup.Label != "-" {
										rowNo += 1
										w.Height(rowNo, s.RowBaseHeight)

										sgLabelCell := s.Cell(s.StartCol, rowNo)
										w.Set(sgLabelCell, fmt.Sprintf("【%s】", *skillGroup.Label))
										w.Merge(sgLabelCell, s.Cell(s.EndCol, rowNo))
										w.CellRangeStyle(sgLabelCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
											s.Alignment(s.HLeftAlignmentIndent3),
											s.Borders(s.SideBorder),
										))

										if skillGroup.Skills != nil {
											builder := util.NewBuilder()
											for idx, skill := range *skillGroup.Skills {
												if idx != 0 {
													builder.Append(", ")
												}
												builder.Append(s.SkillWithVersion(skill))
											}
											if builder.HasError() {
												fmt.Println(builder.ErrorMsg())
												os.Exit(1)
											}

											rowNo += 1
											w.Height(rowNo, s.CalcHeight(builder.ToString()))

											skillNameCell := s.Cell(s.StartCol, rowNo)
											w.Set(skillNameCell, builder.ToString())
											w.Merge(skillNameCell, s.Cell(s.EndCol, rowNo))
											w.CellRangeStyle(skillNameCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
												s.Alignment(s.HLeftAlignmentIndent4),
												s.Borders(s.SideBorder),
											))
										}
									}
								}
							}

							/*
							 * 空行
							 */
							rowNo += 1
							w.Height(rowNo, s.RowBaseHeight)

							spaceCell := s.Cell(s.StartCol, rowNo)
							w.Set(spaceCell, "")
							w.Merge(spaceCell, s.Cell(s.EndCol, rowNo))
							w.CellRangeStyle(spaceCell, s.Cell(s.EndCol, rowNo), s.NewStyle(
								s.Borders(s.SideBottomBorder),
							))
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
