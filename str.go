package str

//NumberOfCharacters counts the number of characters in a string and returns it
func NumberOfCharacters(s string) int {
	c := 0
	for range s {
		c++
	}
	return c
}
