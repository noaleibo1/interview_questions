package main

import (
	"go/types"
	"math/rand"
)

func main() {

}

func fibonaciReq(n int) int {
	if n <= 2 {
		return 1
	}

	return fibonaciReq(n-1) + fibonaciReq(n-2)
}

func fibonaciIter(n int) int {
	previous1 := 1
	previous2 := 1
	for i := 2; i < n; i++ {
		temp := previous1 + previous2
		previous1 = previous2
		previous2 = temp
	}

	return previous2
}

// 1 -> 1
// 2 -> 1
// 3 -> 2
// 4 -> 3
// 5 -> 5

// []()} -> false
// [(]) -> false

func isValidParanteces(str string) bool {
	var mapLeftParToRight map[string]string
	var mapRightToLeft map[string]string

	mapLeftParToRight["("] = ")"
	mapLeftParToRight["["] = "]"
	mapLeftParToRight["{"] = "}"

	mapRightToLeft[")"] = "("
	mapRightToLeft["]"] = "["
	mapRightToLeft["}"] = "{"

	var arr []string
	for i := 0; i < len(str); i++ {
		if _, ok := mapLeftParToRight[string(str[i])]; ok == true {
			arr = append(arr, string(str[i]))
		} else if _, ok := mapRightToLeft[string(str[i])]; ok == true {
			if len(arr) == 0 {
				return false
			}
			if arr[len(arr)-1] == string(str[i]) {
				arr = arr[:len(arr)-1]
			} else {
				return false
			}
		}
	}
	if len(arr) > 0 {
		return false
	}
	return true
}

func getRightMatch(str string) string {
	var res string
	switch str {
	case "(":
		res = ")"
	case "[":
		res = "]"
	case "{":
		res = "}"
	}
	return res
}

func isLeftPar(str string) bool {
	return str == string('[') || str == string('(') || str == string('{')
}

func isRightPar(str string) bool {
	return str == string(']') || str == string(')') || str == string('}')
}

// [1,2,3,4]
// 1 -> 2 , 2 -> 3, 3 -> 4, 4 -> 1

func GamadAnak(arr []int) map[int]int {
	var result map[int]int
	var arrResult []int

	for i := len(arr); i > 0; i-- {
		randomNumberGiver := rand.Intn(len(arr))
		arrResult = append(arr, arr[randomNumberGiver])
		arrTemp := arr[:i]
		arrTempRight := arr[i+1:]
		arr = arrTemp + arrTempRight
	}

	for i := 0; i < len(arrResult); i++ {
		result[arrResult[i]] = arrResult[i+1%len(arrResult)]
	}
	return result
}

// [4,2,1,3]

// [[1,3],[3,4]], [2,5] -> true | [1:1,2:1,3:1,4:1]
// [[1,3],[2,4]], [2,5] -> false
// [[1,7],[2,3][3,4]] , [5,6] ->
// [[1,2],[1,4],[5,6]], [1,10] -> true

func isValidToRecord(records [][]int, newRecord []int) bool {
	// var recordsMap map[int]int
	counter := 0
	for i, record := range records {
		if i < len(records)-2 && (record[0] < records[i+1][0] && newRecord[0] < record[1]) || (newRecord[1] < record[1] && record[0] < newRecord[1]) ||
			(newRecord[0] < record[0] && record[0] < newRecord[1]) || (record[1] < newRecord[1] && newRecord[0] < record[1]) {
			counter++
			if counter > 2 {
				return false
			}
		}

		//
		// for i := record[0]; i < record[1]; i++ {
		// 	if val, ok := recordsMap[i]; ok == true{
		// 		recordsMap[i] = val+1
		// 	} else {
		// 		recordsMap[i] = 1
		// 	}
		// }
	}
	// for i := newRecord[0]; i <= newRecord[1]; i++ {
	// 	if val, ok := recordsMap[i]; ok == true && val == 2{
	// 		return false
	// 	}
	// }

	return true
}
