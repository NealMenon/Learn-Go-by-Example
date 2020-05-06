package iteration

// Repeat character n times
func Repeat(character string, n int) string {
	var repeatedString string
	for i := 0; i < n; i++ {
		repeatedString += character
	}
	return repeatedString
}
