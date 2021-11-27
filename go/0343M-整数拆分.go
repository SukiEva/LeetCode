/*
 * @Author: SukiEva
 * @Date: 2021-11-27 14:59:53
 * @Description: https://github.com/SukiEva
 * わたし、気になります！
 */

package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
