package median_of_three

import "fmt"

func secondIsMedian(input []int) bool {
	return firstIsMedian([]int{input[1], input[0], input[2]})
}

func firstIsMedian(input []int) bool {
	return checkIfMedian(input[0], input[1:])
}

func checkIfMedian(i int, array []int) bool {
	return (i >= array[0] && i <= array[1]) || (i <= array[0] && i >= array[1])
}

func naiiveSolution(input []int) {
	if firstIsMedian(input) {
		fmt.Println(input[0])
	} else if secondIsMedian(input) {
		fmt.Println(input[1])
	} else {
		fmt.Println(input[2])
	}
}
