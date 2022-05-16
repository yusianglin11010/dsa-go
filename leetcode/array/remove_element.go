package main

import "fmt"

func main() {
	test_arr := [][]int{{3, 2, 2, 3}, {0, 1, 2, 2, 3, 0, 4, 2}}
	test_val := []int{3, 2}
	for i := 0; i < len(test_arr); i++ {
		fmt.Println(removeElement(test_arr[i], test_val[i]))
	}
}

func removeElement(nums []int, val int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	i := 0
	for key, v := range m {
		if key != val {
			for v != 0 {
				nums[i] = key
				v--
				i++
			}
		}
	}
	return len(nums) - m[val]
}
