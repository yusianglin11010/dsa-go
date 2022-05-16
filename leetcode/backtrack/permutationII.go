package main

import "fmt"

func main() {
	test := [][]int{{1, 2, 3, 4}, {1, 1, 2}}
	for _, t := range test {
		fmt.Println("The all possible permutation of", t, "is:", permuteUnique(t))
	}
}

func permuteUnique(nums []int) [][]int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	ans := [][]int{}
	backtrack([]int{}, &m, &ans, len(nums))
	return ans
}
func backtrack(tmp []int, m *map[int]int, ans *[][]int, length int) {
	if len(tmp) == length {
		append_tmp := []int{}
		append_tmp = append(append_tmp, tmp...)
		*ans = append(*ans, append_tmp)
	}
	for num := range *m {
		if val, _ := (*m)[num]; val > 0 {
			tmp = append(tmp, num)
			(*m)[num]--
			backtrack(tmp, m, ans, length)
			(*m)[num]++
			tmp = tmp[:len(tmp)-1]
		}
	}
}
