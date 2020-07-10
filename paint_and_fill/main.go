package main

import "fmt"

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 0}, {1, 0, 0}}
	fmt.Println("Before paint")
	printMatrix(matrix)

	paintedMatrix := fillAndPaint(matrix, 1, 2, 2)
	fmt.Println("After paint")
	printMatrix(paintedMatrix)
}

func fillAndPaintReq(matrix [][]int, x int, y int, color int, colorToPaint int) [][]int {
	if matrix[x][y] == color {
		return matrix
	} else if matrix[x][y] == colorToPaint {
		matrix[x][y] = color
		if x+1 < len(matrix) {
			fillAndPaintReq(matrix, x+1, y, color, colorToPaint)
		}
		if x-1 >= 0 {
			fillAndPaintReq(matrix, x-1, y, color, colorToPaint)
		}
		if y+1 < len(matrix[0]) {
			fillAndPaintReq(matrix, x, y+1, color, colorToPaint)
		}
		if y-1 >= 0 {
			fillAndPaintReq(matrix, x, y-1, color, colorToPaint)
		}
	}

	return matrix
}

func fillAndPaint(matrix [][]int, x int, y int, color int) [][]int{
	return fillAndPaintReq(matrix, x,y,color, matrix[x][y])
}


func printMatrix(matrix [][]int) {
	for _, row := range matrix{
		fmt.Println(row)
	}
}