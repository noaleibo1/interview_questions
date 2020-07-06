package theoretical_numbers

import "fmt"

func main() {
	number := getNthNumber(6)
	fmt.Printf("N-th number is: %s", number)
}

func getNthNumber(n int) (nthNumber string) {
	currentNumber := "1"
	for i := 0; i < n; i++ {
		currentNumber = getNextNumber(currentNumber)
	}
	nthNumber = currentNumber
	return
}

func getNextNumber(currentNumber string) (nextNumber string) {
	counter := 1
	for i := 0; i < len(currentNumber); i++ {
		if i == len(currentNumber)-1 {
			nextNumber = fmt.Sprintf("%s%d%s", nextNumber, counter, string(currentNumber[i]))
			break
		}
		if currentNumber[i] == currentNumber[i+1] {
			counter++
		} else {
			nextNumber = fmt.Sprintf("%s%d%s", nextNumber, counter, string(currentNumber[i]))
			counter = 1
		}
	}

	return
}
