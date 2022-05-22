package main

import "fmt"

func main() {
	test := [][2]int{{3, 5}, {3, 7}, {1, 1}}
	for _, t := range test {
		fmt.Println(uniquePaths(t[0], t[1]))
	}
}

func uniquePaths(m int, n int) int {
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, 1)
	}
	for i := 0; i < m-1; i++ {
		for j := 1; j < n; j++ {
			arr[j] = arr[j] + arr[j-1]
		}
	}
	return arr[n-1]
}
