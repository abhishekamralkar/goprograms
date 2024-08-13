package dectobase

import "strings"

func ConvertToBase(n, base int) string {
	if base < 2 || base > 36 {
		return "Base out of range"
	}

	const chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder

	for n > 0 {
		rem := n % base
		result.WriteString(string(chars[rem]))
		n /= base
	}

	return reverseString(result.String())
}

func reverseString(s string) string {
	var res string
	for _, i := range s {
		res = string(i) + res
	}

	return res
}
