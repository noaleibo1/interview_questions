package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getAllSubgroups([]int{1, 2, 3}))
}

func getAllSubgroups(group []int) (subgroups [][]int) {
	binaryArrays := getAllBinaryNumbersOfLengthN(len(group))
	for _, binaryNumber := range binaryArrays {
		var subgroup []int
		for i, bit := range binaryNumber {
			if bit == '1' {
				subgroup = append(subgroup, group[i])
			}
		}
		subgroups = append(subgroups, subgroup)
	}
	return
}

func getAllBinaryNumbersOfLengthN(n int) (binaryNumbers []string) {
	maximumNumber := int(math.Pow(2, float64(n)))
	for i := 0; i < maximumNumber; i++ {
		bit := fmt.Sprintf("%b", i)
		for len(bit) < n {
			bit = fmt.Sprintf("%c%s", '0', bit)
		}
		fmt.Println(bit)
		binaryNumbers = append(binaryNumbers, bit)
	}
	return
}
