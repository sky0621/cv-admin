package submission

import (
	"errors"
	"fmt"
	"github.com/sky0621/cv-admin/src/rest"
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
