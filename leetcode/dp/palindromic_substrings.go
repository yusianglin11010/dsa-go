package main

import "fmt"

func main() {
	test := []string{"aaa", "abc", "abcba", "abccba"}
	for _, t := range test {
		fmt.Println("The string", t, "has", countSubstrings(t), "palindrom substrings")
	}
}

func countSubstrings(s string) int {
	ans, right := 0, len(s)
	for i := 0; i < right; i++ {
		for _, tmp := range [][]int{{i, i}, {i, i + 1}} {
			a, b := tmp[0], tmp[1]
			for a >= 0 && b < right && s[a] == s[b] {
				a--
				b++
			}
			ans += (b - a) / 2
		}
	}
	return ans
}
