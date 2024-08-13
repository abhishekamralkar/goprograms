package factorial

func Fac(n int) (result int) {
	if n > 0 {
		result = n * Fac(n-1)
		return result
	}
	return 1
}
