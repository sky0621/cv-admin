package submission

import (
	"github.com/xuri/excelize/v2"
)

// TODO: １シートのみ使うことを前提としている。複数シート使う場合は再検討。

type ExcelizeWrapperOptions struct {
	SheetName string
}

type ExcelizeWrapperOption func(file *excelize.File)

func SheetName(n string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		f.SetSheetName("Sheet1", n)
	}
}

func PaperSize(s int) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		f.SetPageLayout(getDefaultSheetName(f), excelize.PageLayoutPaperSize(s))
	}
}

func SheetPassword(p string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.ProtectSheet(getDefaultSheetName(f), &excelize.FormatSheetProtection{
			AlgorithmName: "SHA-512",
			Password:      p,
			EditScenarios: false,
		}); err != nil {
			panic(err)
		}
	}
}

func NewExcelizeWrapper(options ...ExcelizeWrapperOption) Wrapper {
	f := excelize.NewFile()
	for _, option := range options {
		option(f)
	}
	return &wrapper{f: f}
}

type Wrapper interface {
	Set(position string, val interface{})
	Merge(from, to string)
	Height(row int, h float64)
	Width(start, end string, wd float64)
	Text(cell string, settings []excelize.RichTextRun)
	CenterStyleID() int
	RightStyleID() int
	LeftStyleID() int
	CellStyle(start, end string, styleID int)

	SaveAs(name string)
}

type wrapper struct {
	f *excelize.File
}

func getDefaultSheetName(f *excelize.File) string {
	return f.GetSheetName(0)
}

func (w *wrapper) Set(position string, val interface{}) {
	if err := w.f.SetCellValue(getDefaultSheetName(w.f), position, val); err != nil {
		panic(err)
	}
}

func (w *wrapper) Merge(from, to string) {
	if err := w.f.MergeCell(getDefaultSheetName(w.f), from, to); err != nil {
		panic(err)
	}
}

func (w *wrapper) Height(row int, h float64) {
	if err := w.f.SetRowHeight(getDefaultSheetName(w.f), row, h); err != nil {
		panic(err)
	}
}

func (w *wrapper) Width(start, end string, wd float64) {
	if err := w.f.SetColWidth(getDefaultSheetName(w.f), start, end, wd); err != nil {
		panic(err)
	}
}

func (w *wrapper) Text(cell string, settings []excelize.RichTextRun) {
	if err := w.f.SetCellRichText(getDefaultSheetName(w.f), cell, settings); err != nil {
		panic(err)
	}
}

var vhCenterAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "center",
}

var hRightAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "right",
}

var hLeftAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "left",
}

func (w *wrapper) getStyleID(a *excelize.Alignment) int {
	styleID, err := w.f.NewStyle(&excelize.Style{Alignment: a})
	if err != nil {
		panic(err)
	}
	return styleID
}

func (w *wrapper) CenterStyleID() int {
	return w.getStyleID(vhCenterAlignment)
}

func (w *wrapper) RightStyleID() int {
	return w.getStyleID(hRightAlignment)
}

func (w *wrapper) LeftStyleID() int {
	return w.getStyleID(hLeftAlignment)
}

func (w *wrapper) CellStyle(start, end string, styleID int) {
	if err := w.f.SetCellStyle(getDefaultSheetName(w.f), start, end, styleID); err != nil {
		panic(err)
	}
}

func (w *wrapper) SaveAs(name string) {
	if err := w.f.SaveAs(name); err != nil {
		panic(err)
	}
}
