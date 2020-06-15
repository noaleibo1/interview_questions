package medean_of_three

import "fmt"

func sortSolution(input []int) {
	sortedInput := mergeSort(input, 0, len(input)-1)

	fmt.Println(sortedInput)
	fmt.Printf("Median is %d", sortedInput[len(sortedInput)/2])
}

func mergeSort(arr []int, i int, j int) []int {
	m := i + (j-i)/2

	if i < j {
		a1 := mergeSort(arr, i, m)
		a2 := mergeSort(arr, m+1, j)
		return merge(a1, a2)
	}
	return []int{arr[i]}
}

func merge(arr1 []int, arr2 []int) (mergedArr []int) {
	i := 0
	j := 0
	for j < len(arr2) && i < len(arr1) {
		if arr1[i] < arr2[j] {
			mergedArr = append(mergedArr, arr1[i])
			i++
		} else {
			mergedArr = append(mergedArr, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		mergedArr = append(mergedArr, arr1[i])
		i++
	}
	for j < len(arr2) {
		mergedArr = append(mergedArr, arr2[j])
		j++
	}
	return
}
