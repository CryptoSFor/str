package str

func ReverseString(s string) string {
	r := []rune(s)
	for a, b := 0, len(r)-1; a < b; a, b = a+1, b-1 {
		r[a], r[b] = r[b], r[a]
	}
	return string(r)
}
