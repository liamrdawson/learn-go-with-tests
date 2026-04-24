package main

func Sum(arr []int) (sum int) {

	for _, val := range arr {
		sum += val
	}

	return sum
}

var a = [3]int{1, 2, 3}

var aCopy = a

func SumAll(numbersToSum ...[]int) []int {

	var sums []int
	for _, val := range numbersToSum {
		sums = append(sums, Sum(val))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, val := range numbersToSum {
		if len(val) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(val[1:]))
		}
	}

	return sums
}
