package submission

import "github.com/xuri/excelize/v2"

const (
	borderStyleDot  = 3
	borderBaseColor = "6495ED"
)

var Border = []excelize.Border{
	{Type: "left", Color: borderBaseColor, Style: borderStyleDot},
	{Type: "top", Color: borderBaseColor, Style: borderStyleDot},
	{Type: "right", Color: borderBaseColor, Style: borderStyleDot},
	{Type: "bottom", Color: borderBaseColor, Style: borderStyleDot},
}

var VhCenterAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "center",
}

var HRightAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "right",
}

var HLeftAlignment = &excelize.Alignment{
	Vertical:   "center",
	Horizontal: "left",
}

var SheetTitleFont = &excelize.Font{
	Size: 16,
	Bold: true,
}

var LastUpdatedFont = &excelize.Font{
	Size: 9,
}

var NameFont = &excelize.Font{
	Size: 18,
	Bold: true,
}

type ExcelizeStyleOption func(style *excelize.Style)

func Alignment(a *excelize.Alignment) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Alignment = a
	}
}

func Borders(b []excelize.Border) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Border = b
	}
}

func Font(f *excelize.Font) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Font = f
	}
}

func Fill(f excelize.Fill) ExcelizeStyleOption {
	return func(style *excelize.Style) {
		style.Fill = f
	}
}

func NewStyle(options ...ExcelizeStyleOption) *excelize.Style {
	s := &excelize.Style{}
	for _, option := range options {
		option(s)
	}
	return s
}
