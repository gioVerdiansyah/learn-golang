package my_modules

func Total(val []int) int {
	var total int
	for _, v := range val {
		total += v
	}

	return total
}

func Avg(val []int) int {
	return Total(val) / len(val)
}

func Min(val []int) int {
	if len(val) == 0 {
		return 0
	}

	var minVal = val[0]
	for _, v := range val {
		if v < minVal {
			minVal = v
		}
	}

	return minVal
}

func Max(val []int) int {
	var maxVal int
	for _, v := range val {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
