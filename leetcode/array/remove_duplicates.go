package main

import "fmt"

func main() {
	nums := [][]int{{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, {0, 1, 2, 3, 3, 4}}
	for _, num := range nums {
		fmt.Println(removeDuplicates(num))
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	left, p := 0, 0
	for left < len(nums)-1 {
		i := 1
		if nums[left] == nums[left+1] {
			for nums[left] == nums[left+i] && left+i < len(nums)-1 {
				i++
			}
		}
		nums[p] = nums[left+i-1]
		p++
		left += i
	}
	if nums[left] != nums[left-1] {
		nums[p] = nums[left]
		p++
	}
	return p
}

func removeDuplicatesExtraSpace(nums []int) int {
	m := make(map[int]bool)
	p := 0
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = true
			nums[p] = num
			p++
		}
	}
	return p
}
