package main

import "fmt"

func main() {
	sum := recursionMultiply(3, 5)
	fmt.Println(sum)
}

func naiveRecursiveMultiply(i int, j int) (sum int) {
	if j == 1 {
		return i
	} else {
		return i + naiveRecursiveMultiply(i, j-1)
	}
}

func recursionMultiply(i int, j int) (sum int) {
	multiplyCache := map[int]int{}
	if i == 0 || j == 0 {
		sum = 0
		return
	} else if j == 1 {
		multiplyCache[1] = i
		sum = i
		return
	} else if v, ok := multiplyCache[j]; ok {
		sum = v
		return
	} else if v, ok := multiplyCache[j/2]; ok {
		if j%2 == 0 {
			sum = v + v
		} else {
			sum = v + v - 1
		}
	} else {
		if j%2 == 0 {
			sum = recursionMultiply(i, j/2) + recursionMultiply(i, j/2)
		} else {
			sum = recursionMultiply(i, j/2) + recursionMultiply(i, j/2) - 1
		}
	}
	return
}
