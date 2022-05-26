package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	var rune_out []rune
	for i := range runes {
		rune_out = append(rune_out, runes[len(runes)-i-1])
	}
	output = string(rune_out)
	return output
}
