package main

import "fmt"

func main() {
	arr := []int{2, 1, 0, 1}
	b := isJumpingPossible(arr)
	fmt.Printf("%v", b)
}

func isJumpingPossible(arr []int) (isPossible bool) {
	if len(arr) > 1 && arr[0] == 0 {
		isPossible = false
		return
	} else if len(arr) == 1 && arr[0] > 0 {
		isPossible = true
		return
	}
	for i := 1; i <= arr[0]; i++ {
		if isJumpingPossible(arr[i:]) {
			isPossible = true
			return
		}
	}
	isPossible = false
	return
}
