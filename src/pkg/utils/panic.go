package utils

func PanicIfNotNil(v any, s string) {
	if v != nil {
		panic(s)
	}
}
