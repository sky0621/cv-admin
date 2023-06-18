package converter

func ToPtr[T any](v T) *T {
	return &v
}

func ToVal[T any](p *T) T {
	return *p
}
