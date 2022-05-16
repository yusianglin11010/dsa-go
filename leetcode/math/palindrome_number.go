package main

import "fmt"

func main() {
	test := []int{121, 1221, 1345, -5, 0}
	for _, t := range test {
		fmt.Println(t, "is palindrome?", isPalindrome(t))
	}
	for _, t := range test {
		fmt.Println(t, "is palindrome?", isPalindromeOptimize(t))
	}
}

// O(n)
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else {
		num := x
		tmp := 0
		for x > 0 {
			tmp = tmp*10 + x%10
			x = x / 10
		}
		return num == tmp
	}
}

// O(n/2)
func isPalindromeOptimize(x int) bool {
	if x < 0 || (x > 0 && x%10 != 0) {
		return false
	} else {
		reverseX := 0
		for x > reverseX {
			reverseX = reverseX*10 + x%10
			x /= 10
		}
		return reverseX == x || reverseX/10 == x
	}
}
