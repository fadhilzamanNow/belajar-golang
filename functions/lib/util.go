package utilfunc

func Substract(a int, b int) int {
	return a + b
}

func ReturnTwo(a, b int) (int, int) {
	return a, b
}

func MultipleUltimate(a ...int) int {
	sum := 1
	for i := range a {
		sum *= a[i]
	}
	return sum
}
