package arrays

func Sum(numbers []int) int {
	// var sum int
	sum := 0
	// for i := 0; i < len(numbers); i++ {
	// 	sum += numbers[i]
	// }
	for _, number := range numbers {
		sum += number
	}
	return sum
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
