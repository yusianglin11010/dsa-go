package main

import "fmt"

func main() {
	test := [][][]int{{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, {{1}}}
	for _, t := range test {
		fmt.Println(longestIncreasingPath(t))
	}

}

func longestIncreasingPath(matrix [][]int) int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	pos := make(map[[2]int]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			pos[[2]int{i, j}] = 0
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			dfs(matrix, i, j, dirs, &pos)
		}
	}
	ans := 0
	for _, v := range pos {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func dfs(matrix [][]int, i int, j int, dirs [][]int, pos *map[[2]int]int) {
	if (*pos)[[2]int{i, j}] > 0 {
		return
	}
	max_val := 0
	(*pos)[[2]int{i, j}] = 1
	for _, dir := range dirs {
		i_2, j_2 := i+dir[0], j+dir[1]
		if i_2 >= 0 && j_2 >= 0 && i_2 < len(matrix) && j_2 < len(matrix[0]) {
			if matrix[i_2][j_2] > matrix[i][j] {
				dfs(matrix, i_2, j_2, dirs, pos)
				if (*pos)[[2]int{i_2, j_2}] > max_val {
					max_val = (*pos)[[2]int{i_2, j_2}]
				}
			}
		}
	}
	(*pos)[[2]int{i, j}] += max_val
}
