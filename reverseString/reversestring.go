package reverseString

func ReverseIt(s string) (result string) {
	for _, v := range s {
		result = string(v) + result

	}
	return
}
