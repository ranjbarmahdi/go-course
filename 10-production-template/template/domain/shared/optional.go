package shared

func CloneOptionalString(v *string) *string {
	if v == nil || *v == "" {
		return nil
	}
	s := *v
	return &s
}

func CloneOptionalFloat64(v *float64) *float64 {
	if v == nil {
		return nil
	}
	f := *v
	return &f
}

func ClonePtr[T any](v *T) *T {
	if v == nil {
		return nil
	}
	c := *v
	return &c
}
