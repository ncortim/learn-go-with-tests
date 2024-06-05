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
