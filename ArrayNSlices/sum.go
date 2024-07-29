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
