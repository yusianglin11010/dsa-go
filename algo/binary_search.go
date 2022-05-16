package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	target := 7
	fmt.Println("Target at ", bs(array, target), "index")
	// Target at  6 index
	array = []int{1, 7, 7, 7, 8}
	fmt.Println("Target at ", bs(array, target), "index")
	// Target at 6 index
	// Can not get target at left side or target at right side
}

// Find element index in a sorted array, if not found, return -1 instead
func bs(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = left + 1
		}
	}
	return -1
}
