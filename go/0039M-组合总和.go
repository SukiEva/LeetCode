/*
 * @Author: SukiEva
 * @Date: 2021-11-13 08:53:07
 * @Description: https://github.com/SukiEva
 * わたし、気になります！
 */
package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	vis := make([]int, 0)
	var dfs func(index, sum int)
	dfs = func(index, sum int) {
		if sum == target {
			tmp := make([]int, len(vis))
			copy(tmp, vis)
			ans = append(ans, tmp)
			return
		}
		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			vis = append(vis, candidates[i])
			dfs(i, sum+candidates[i]) // 重复读取不＋1
			vis = vis[:len(vis)-1]
		}
	}
	sort.Ints(candidates)
	dfs(0, 0)
	return ans
}
