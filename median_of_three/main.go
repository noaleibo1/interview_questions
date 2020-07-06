package median_of_three

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	intInput, err := covertToInt(input)
	if err != nil {
		panic(err.Error())
	}
	// if !validInput(intInput) {
	// 	panic("Input should be 3 integers")
	// }

	sortSolution(intInput)
}

func covertToInt(input []string) ([]int, error) {
	var intArray []int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("some of the input is invalid")
		}
		intArray = append(intArray, i)
	}
	return intArray, nil
}

func validInput(input []int) bool {
	if len(input) != 3 {
		return false
	}
	return true
}
