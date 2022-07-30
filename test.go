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
