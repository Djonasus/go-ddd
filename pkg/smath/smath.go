package smath

func Clamp(t int, a int, b int) int {
	if t > b {
		return a
	} else if t < a {
		return b
	} else {
		return t
	}
}

const (
	Huge int = 2147483647
)
