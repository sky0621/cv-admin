package submission

import (
	"errors"
	"fmt"
	"github.com/sky0621/cv-admin/src/rest"
	"time"
)

func Cell(col string, row int) string {
	return fmt.Sprintf("%s%d", col, row)
}

func QualificationName(org *rest.QualificationOrg, name *rest.QualificationName) string {
	if name == nil {
		panic(errors.New("no qualification name"))
	}
	if org == nil {
		return *name
	}
	return fmt.Sprintf("[%s] %s", *org, *name)
}

func QualificationGotDate(d *rest.QualificationGotDate) string {
	if d == nil {
		return "-"
	}
	return d.Format("2006-01-02")
}

func URL(u *rest.Url) string {
	if u == nil {
		return "-"
	}
	return *u
}

func CareerPeriod(career rest.UserCareer) string {
	f := career.From
	t := career.To
	if f == nil || t == nil {
		return ""
	}
	dms := MergeMonths(DiffMonths(*f.Year, *f.Month, *t.Year, *t.Month))

	return fmt.Sprintf("%d / %d - %d / %d ( %s )", *f.Year, *f.Month, *t.Year, *t.Month, dms)
}

func CareerGroupPeriod(careerGroup rest.UserCareerGroup) string {
	if careerGroup.Careers == nil {
		return ""
	}

	var startYear, startMonth, endYear, endMonth int
	for _, career := range *careerGroup.Careers {
		if career.From == nil || career.To == nil {
			continue
		}

		if startYear == 0 {
			startYear = *career.From.Year
			startMonth = *career.From.Month
		} else if startYear == *career.From.Year {
			if startMonth > *career.From.Month {
				startYear = *career.From.Year
				startMonth = *career.From.Month
			}
		} else if startYear > *career.From.Year {
			startYear = *career.From.Year
			startMonth = *career.From.Month
		}

		if endYear == *career.To.Year {
			if endMonth < *career.To.Month {
				endYear = *career.To.Year
				endMonth = *career.To.Month
			}
		} else if endYear < *career.To.Year {
			endYear = *career.To.Year
			endMonth = *career.To.Month
		}
	}

	if startYear == 0 || endYear == 0 {
		return ""
	}

	dms := MergeMonths(DiffMonths(startYear, startMonth, endYear, endMonth))

	return fmt.Sprintf("%d / %d - %d / %d ( %s )", startYear, startMonth, endYear, endMonth, dms)
}

func Age(birthYear, birthMonth, birthDay int, now time.Time) int {
	age := now.Year() - birthYear

	if int(now.Month()) > birthMonth {
		return age
	}

	if int(now.Month()) == birthMonth && now.Day() >= birthDay {
		return age
	}

	return age - 1
}

func DiffMonths(fromYear, fromMonth, toYear, toMonth int) int {
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

func MergeMonths(dms int) string {
	if dms < 12 {
		return fmt.Sprintf("%dヶ月", dms)
	}

	year := dms / 12
	remain := dms % 12
	if remain == 0 {
		return fmt.Sprintf("%d年", year)
	}

	return fmt.Sprintf("%d年%dヶ月", year, remain)
}

// TODO: ２行を超える場合も対応！
func SkillWithVersion(careerSkill rest.CareerSkill) string {
	if careerSkill.Skill == nil {
		return "-"
	}
	if careerSkill.Version == nil {
		return *careerSkill.Skill.Name
	}
	return fmt.Sprintf("%s(%s)", *careerSkill.Skill.Name, *careerSkill.Version)
}

func CalcHeight(str string) float64 {
	if len(str) > BaseColWidth*(ViewAreaColNum-1) {
		return RowBaseHeight * 1.5
	}
	return RowBaseHeight
}
