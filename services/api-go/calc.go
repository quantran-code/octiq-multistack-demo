package api

func Add(left, right int) int {
	return left + right
}

func Divide(left, right int) (float64, bool) {
	if right == 0 {
		return 0, false
	}
	return float64(left) / float64(right), true
}
