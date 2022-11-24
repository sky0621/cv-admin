package submission

import (
	"github.com/xuri/excelize/v2"
)

func NewExcelizeWrapper(sheetName string, worksheetSize int) Wrapper {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	f.SetPageLayout(sheetName, excelize.PageLayoutPaperSize(worksheetSize))
	return &wrapper{f: f, sheetName: sheetName}
}

type Wrapper interface {
	Set(position string, val interface{})
	Merge(from, to string)
	Height(row int, h float64)
	Width(start, end string, wd float64)
	Text(cell string, settings []excelize.RichTextRun)
	CenterStyleID() int
	RightStyleID() int
	CellStyle(start, end string, styleID int)

	SaveAs(name string)
}

type wrapper struct {
	f         *excelize.File
	sheetName string
}

func (w *wrapper) Set(position string, val interface{}) {
	if err := w.f.SetCellValue(w.sheetName, position, val); err != nil {
		panic(err)
	}
}

func (w *wrapper) Merge(from, to string) {
	if err := w.f.MergeCell(w.sheetName, from, to); err != nil {
		panic(err)
	}
}

func (w *wrapper) Height(row int, h float64) {
	if err := w.f.SetRowHeight(w.sheetName, row, h); err != nil {
		panic(err)
	}
}

func (w *wrapper) Width(start, end string, wd float64) {
	if err := w.f.SetColWidth(w.sheetName, start, end, wd); err != nil {
		panic(err)
	}
}

func (w *wrapper) Text(cell string, settings []excelize.RichTextRun) {
	if err := w.f.SetCellRichText(w.sheetName, cell, settings); err != nil {
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

func (w *wrapper) CellStyle(start, end string, styleID int) {
	if err := w.f.SetCellStyle(w.sheetName, start, end, styleID); err != nil {
		panic(err)
	}
}

func (w *wrapper) SaveAs(name string) {
	if err := w.f.SaveAs(name); err != nil {
		panic(err)
	}
}
