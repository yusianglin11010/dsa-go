package main

import "fmt"

func main() {
	array := []int{1, 7, 7, 7, 7, 12}
	target := 7
	fmt.Println("Target at ", bs(array, target), "index")
}

func bs(nums []int, target int) int {
	left, right := 0, len(nums)
	for right > left {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
