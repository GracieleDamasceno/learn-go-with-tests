package main

func SumArraySlices(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllArrays(valuesToSum ...[]int) []int {
	var sumAll []int
	for _, numbers := range valuesToSum {
		sumAll = append(sumAll, SumArraySlices(numbers))
	}
	return sumAll
}

func SumAllTails(valuesToSum ...[]int) []int {
	var sumTails []int
	for _, numbers := range valuesToSum {
		if len(numbers) == 0 {
			sumTails = append(sumTails, 0)

		} else {
			tail := numbers[1:]
			sumTails = append(sumTails, SumArraySlices(tail))
		}
	}
	return sumTails
}
