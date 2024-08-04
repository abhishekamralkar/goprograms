package findnumber

func FindNumber(l []int, n int) bool {
	for _, i := range l {
		if i == n {
			return true
		}
	}
	return false
}
