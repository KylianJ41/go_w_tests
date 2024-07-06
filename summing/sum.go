package summing

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	/*lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i := 0; i < lengthOfNumbers; i++ {
		sums[i] = Sum(numbersToSum[i])
	}*/

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
