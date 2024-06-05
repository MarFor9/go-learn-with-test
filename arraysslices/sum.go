package arraysslices

func Sum(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

// Variadic functions - the method signature can work with different number of aguments. It's done by using "..." syntax before type.
// For example: x ...int[] <- this will allow put in to method varying number of slices
func SumAll(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numberToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numberToSum {
		if len(numbers) > 0 {
			sums = append(sums, Sum(numbers[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
