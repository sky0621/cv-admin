package submission

import "github.com/xuri/excelize/v2"

const (
	RowBaseHeight float64 = 25.0
	RowMaxHeight  float64 = 409.0
)

const (
	headerColor     = "9FE2BF"
	borderBaseColor = headerColor
	borderBaseStyle = 5
)

var FullBorder = []excelize.Border{
	{Type: "left", Color: borderBaseColor, Style: borderBaseStyle},
	{Type: "top", Color: borderBaseColor, Style: borderBaseStyle},
	{Type: "right", Color: borderBaseColor, Style: borderBaseStyle},
	{Type: "bottom", Color: borderBaseColor, Style: borderBaseStyle},
}

var LeftBorder = []excelize.Border{
	{Type: "left", Color: borderBaseColor, Style: borderBaseStyle},
}

var BottomBorder = []excelize.Border{
	{Type: "bottom", Color: borderBaseColor, Style: borderBaseStyle},
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

var HeaderFill = excelize.Fill{
	Type:    "pattern",
	Color:   []string{headerColor},
	Pattern: 1,
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
