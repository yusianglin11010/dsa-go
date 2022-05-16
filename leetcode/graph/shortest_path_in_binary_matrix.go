package main

import "fmt"

func main() {
	test := [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}
	fmt.Println(shortestPathBinaryMatrix(test))
}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] > 0 || grid[n-1][n-1] > 0 {
		return -1
	}
	seen := make(map[[2]int]bool)
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	dq := [][3]int{{0, 0, 1}}
	for len(dq) != 0 {
		i, j, d := dq[0][0], dq[0][1], dq[0][2]
		if i == n-1 && j == n-1 {
			return d
		}
		dq = dq[1:]
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			if 0 <= x && n > x && 0 <= y && n > y {
				if _, ok := seen[[2]int{x, y}]; !ok && grid[x][y] == 0 {
					seen[[2]int{x, y}] = true
					dq = append(dq, [3]int{x, y, d + 1})
				}
			}
		}
	}
	return -1
}
