package sumlistnumbers

func SumList(l []int) int {
	total := 0
	for _, n := range l {
		total += n
	}
	return total
}
