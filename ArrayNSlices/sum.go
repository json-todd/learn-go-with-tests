package main

func Sum(numbers []int) (sum int) {
    for _, number := range numbers {
        sum += number
    }
    return
}

func SumAll(numbersToSum...[]int) []int {
    var sumSlices []int
    for _, slize := range numbersToSum {
        sum := Sum(slize)
        sumSlices = append(sumSlices, sum)
    }
    return sumSlices
}

func SumAllTail(numbersToSum...[]int) []int {
	var tailNumbersToSum [][]int	
	for _, slize := range numbersToSum {
		tailSlize := slize[1:]
		tailNumbersToSum = append(tailNumbersToSum, tailSlize)
	}

	return SumAll(tailNumbersToSum...)
}
