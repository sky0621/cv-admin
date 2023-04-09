package submission

import (
	"github.com/xuri/excelize/v2"
)

// TODO: １シートのみ使うことを前提としている。複数シート使う場合は再検討。

type ExcelizeWrapperOption func(file *excelize.File)

func SheetName(n string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		f.SetSheetName("Sheet1", n)
	}
}

func PaperSize(s int) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		f.SetPageLayout(getDefaultSheetName(f), &excelize.PageLayoutOptions{
			Size: &s,
		})
	}
}

func SheetPassword(p string) ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.ProtectSheet(getDefaultSheetName(f), &excelize.SheetProtectionOptions{
			AlgorithmName: "SHA-512",
			Password:      p,
			EditScenarios: false,
		}); err != nil {
			panic(err)
		}
	}
}

func SheetView() ExcelizeWrapperOption {
	return func(f *excelize.File) {
		if err := f.SetSheetView(getDefaultSheetName(f), 0, &excelize.ViewOptions{
			ShowGridLines: func() *bool {
				b := false
				return &b
			}(),
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
	Set(position string, val any)
	Merge(from, to string)
	Height(row int, h float64)
	Width(start, end string, wd float64)
	Text(cell string, settings []excelize.RichTextRun)
	CellStyle(cell string, style *excelize.Style)
	CellRangeStyle(start, end string, style *excelize.Style)
	HeaderCellStyle(cell string)
	HeaderCellRangeStyle(start, end string)
	ValueCellRangeStyle(start, end string)
	CellExternalHyperLink(cell, url string)
	AddPicture(cell, path string)
	AddPictureFromBytes(cell, name, extension string, file []byte)
	InsertPageBreak(cell string)

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

func (w *wrapper) style(s *excelize.Style) int {
	styleID, err := w.f.NewStyle(s)
	if err != nil {
		panic(err)
	}
	return styleID
}

func (w *wrapper) CellStyle(cell string, style *excelize.Style) {
	if err := w.f.SetCellStyle(getDefaultSheetName(w.f), cell, cell, w.style(style)); err != nil {
		panic(err)
	}
}
func (w *wrapper) CellRangeStyle(start, end string, style *excelize.Style) {
	if err := w.f.SetCellStyle(getDefaultSheetName(w.f), start, end, w.style(style)); err != nil {
		panic(err)
	}
}

func (w *wrapper) HeaderCellStyle(cell string) {
	s := NewStyle(
		Alignment(HLeftAlignment),
		Borders(FullBorder),
		Fill(HeaderFill),
	)
	w.CellStyle(cell, s)
}

func (w *wrapper) HeaderCellRangeStyle(start, end string) {
	s := NewStyle(
		Alignment(HLeftAlignment),
		Borders(FullBorder),
		Fill(HeaderFill),
	)
	w.CellRangeStyle(start, end, s)
}

func (w *wrapper) ValueCellRangeStyle(start, end string) {
	s := NewStyle(
		Alignment(HLeftAlignment),
		Borders(FullBorder),
	)
	w.CellRangeStyle(start, end, s)
}

func (w *wrapper) CellExternalHyperLink(cell, url string) {
	if err := w.f.SetCellHyperLink(getDefaultSheetName(w.f), cell, url, "External", excelize.HyperlinkOpts{
		Display: &url,
	}); err != nil {
		panic(err)
	}
}

func (w *wrapper) AddPicture(cell, path string) {
	if err := w.f.AddPicture(getDefaultSheetName(w.f), cell, path, nil); err != nil {
		panic(err)
	}
}

func (w *wrapper) AddPictureFromBytes(cell, name, extension string, file []byte) {
	if err := w.f.AddPictureFromBytes(getDefaultSheetName(w.f), cell, name, extension, file, &excelize.GraphicOptions{
		OffsetX: 16, OffsetY: 16,
	}); err != nil {
		panic(err)
	}
}

func (w *wrapper) InsertPageBreak(cell string) {
	if err := w.f.InsertPageBreak(getDefaultSheetName(w.f), cell); err != nil {
		panic(err)
	}
}

func (w *wrapper) SaveAs(name string) {
	if err := w.f.SaveAs(name); err != nil {
		panic(err)
	}
}
