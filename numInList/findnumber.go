package findnumber

import (
	"errors"
)

func FindNumber(l []int, n int) (bool, error) {
	if len(l) < 1 {
		return false, errors.New("empty data")
	}
	for _, i := range l {
		if i == n {
			return true, nil
		}
	}
	return false, nil
}
