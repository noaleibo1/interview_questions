package main

import (
	"fmt"
)

func main() {
	printMatrix()
	path := getPathInMaze(maze[1][0])
	printPath(path)
}

var maze = [][]*cell{{{i: 0, j: 0, isClear: false, isVisited: false}, {i: 0, j: 1, isClear: false, isVisited: false}, {i: 0, j: 2, isClear: false, isVisited: false}, {i: 0, j: 3, isClear: false, isVisited: false}},
	{{i: 1, j: 0, isClear: true, isVisited: false}, {i: 1, j: 1, isClear: true, isVisited: false}, {i: 1, j: 2, isClear: true, isVisited: false}, {i: 1, j: 3, isClear: false, isVisited: false}},
	{{i: 2, j: 0, isClear: false, isVisited: false}, {i: 2, j: 1, isClear: false, isVisited: false}, {i: 2, j: 2, isClear: true, isVisited: false}, {i: 2, j: 3, isClear: true, isVisited: false}},
	{{i: 3, j: 0, isClear: false, isVisited: false}, {i: 3, j: 1, isClear: false, isVisited: false}, {i: 3, j: 2, isClear: false, isVisited: false}, {i: 3, j: 3, isClear: false, isVisited: false}},
}

type cell struct {
	i         int
	j         int
	isClear   bool
	isVisited bool
}

func (c cell) isExit() bool {
	if c.i == len(maze)-1 || c.i == 0 || c.j == 0 || c.j == len(maze[0])-1 {
		return true
	}
	return false
}

func getPathInMaze(startCell *cell) (path []*cell) {
	path = append(path, startCell)
	startCell.isVisited = true
	var neighbor *cell
	for len(path) > 0 {
		neighbor = getClearUnvisitedNeighbor(path[len(path)-1])
		if neighbor != nil {
			path = append(path, neighbor)
			neighbor.isVisited = true
			if neighbor.isExit() {
				return
			}
		} else {
			path = path[:len(path)-1] // Remove last step from path
		}
	}
	return
}

func getClearUnvisitedNeighbor(c *cell) (neighbor *cell) {
	if c.i < len(maze)-1 && maze[c.i+1][c.j].isClear && (maze[c.i+1][c.j].isVisited == false) { // Up neighbor
		neighbor = maze[c.i+1][c.j]
	} else if c.i > 0 && maze[c.i-1][c.j].isClear && (maze[c.i-1][c.j].isVisited == false) { // Down neighbor
		neighbor = maze[c.i-1][c.j]
	} else if c.j > 0 && maze[c.i][c.j-1].isClear && (maze[c.i][c.j-1].isVisited == false) { // Left neighbor
		neighbor = maze[c.i][c.j-1]
	} else if c.j < len(maze[0])-1 && maze[c.i][c.j+1].isClear && (maze[c.i][c.j+1].isVisited == false) { // Right neighbor
		neighbor = maze[c.i][c.j+1]
	}
	return
}

func printMatrix() {
	for _, line := range maze {
		var toPrint string
		for _, cell := range line {
			if cell.isClear {
				toPrint = fmt.Sprintf("%s %s", toPrint, "0")
			} else {
				toPrint = fmt.Sprintf("%s %s", toPrint, "1")
			}
		}
		fmt.Println(toPrint)
		toPrint = ""
	}
}

func printPath(path []*cell) {
	for _, c := range path {
		fmt.Printf("(%d,%d)", c.i, c.j)
	}
}
