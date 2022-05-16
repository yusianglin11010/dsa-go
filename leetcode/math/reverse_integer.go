package main

import (
	"fmt"
	"math"
)

func main() {
	test := []int{121, 321, -123, 10001}
	for _, t := range test {
		fmt.Println(t, "reverse is:", reverse(t))
	}
}

func reverse(x int) int {
	result := 0
	state := 1
	if x < 0 {
		state = -1
		x *= state
	}
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result > int(math.Pow(2, 31)-1) || result < int(math.Pow(-2, 31)) {
		return 0
	} else {
		return result * state
	}

}
