package main

import "fmt"

func main() {
	test := [][]int{{1, 2, 3}, {4, 3, 2}}
	for _, t := range test {
		fmt.Println("The all possible permutation of ", t, "is:", permute(t))
	}
}

func permute(nums []int) [][]int {
	ans, tmp := [][]int{}, []int{}
	length := len(nums)
	backtrack(nums, tmp, length, &ans)
	return ans
}

func backtrack(nums []int, tmp []int, length int, ans *([][]int)) {
	if len(tmp) == length {
		append_tmp := []int{}
		append_tmp = append(append_tmp, tmp...)
		*ans = append(*ans, append_tmp)
		return
	}
	for i, n := range nums {
		tmp = append(tmp, n)
		tmp_num := []int{}
		tmp_num = append(tmp_num, nums[:i]...)
		tmp_num = append(tmp_num, nums[i+1:]...)
		backtrack(tmp_num, tmp, length, ans)
		tmp = tmp[:len(tmp)-1]
	}
}
