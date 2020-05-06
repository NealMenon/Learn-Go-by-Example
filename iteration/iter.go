package iteration

// RepeatN character n times
func RepeatN(character string, n int) string {
	if n == 0 {
		n = 5
	}
	var repeatedString string
	for i := 0; i < n; i++ {
		repeatedString += character
	}
	return repeatedString
}

// Repeat5 character 5 times
func Repeat5(character string) string {
	var repeatedString string
	for i := 0; i < 5; i++ {
		repeatedString += character
	}
	return repeatedString
}
