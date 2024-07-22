package functions 

func LargeSum(x ...int) (int, int) {
	var sum int
	for _, v := range x{
		sum += v
	}

	return sum, len(x)
}