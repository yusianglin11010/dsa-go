package main

import "fmt"

func two_sum_ON2(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func two_sum_ON(nums []int, target int) []int {
	d := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := d[target-nums[i]]; ok {
			return []int{d[target-nums[i]], i}
		}
		d[nums[i]] = i
	}
	return nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Println(two_sum_ON(nums, target))
}
