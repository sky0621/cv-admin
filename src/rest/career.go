package rest

// CareerPeriodIf キャリアの開始・終了年月を束ねるためのマーカーインタフェース
type CareerPeriodIf interface {
	Set(year, month *int)
}

func (f *CareerPeriodFrom) Set(year, month *int) {
	f.Year = year
	f.Month = month
}

func (f *CareerPeriodTo) Set(year, month *int) {
	f.Year = year
	f.Month = month
}
