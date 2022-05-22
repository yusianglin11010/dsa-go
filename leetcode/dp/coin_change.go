package main

import "fmt"

func main() {

	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}

func coinChange(coins []int, amount int) int {
	dp := []int{}
	for i := 0; i <= amount; i++ {
		dp = append(dp, amount+1)
	}
	dp[0] = 0
	for i, _ := range dp {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i] > dp[i-coin]+1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
