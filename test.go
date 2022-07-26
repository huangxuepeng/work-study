package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	arr := [][]int{}
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > pre[1] {
			pre = intervals[i]
			arr = append(arr, pre)
		} else {
			pre[1] = max(pre[1], intervals[i][1])
		}
		
	}
	arr = append(arr, pre)
	return arr
}
for i := 1; i < len(intervals); i++ {
	cur := intervals[i]
	if prev[1] < cur[0] { // 没有一点重合
		res = append(res, prev)
		prev = cur
	} else { // 有重合
		prev[1] = max(prev[1], cur[1])
	}
}
