package slices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int
	for _, numbers := range numbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
