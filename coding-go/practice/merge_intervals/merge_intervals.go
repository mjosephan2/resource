package merge_intervals

import "sort"

/**
https://leetcode.com/problems/merge-intervals/description/?envType=study-plan-v2&envId=top-interview-150
**/
func Merge(intervals [][]int) [][]int {
	var res [][]int
	sortIntervals(intervals)
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}
		curInterval := interval
		for len(res) > 0 {
			lastInterval := res[len(res)-1]
			if isOverlap(lastInterval, curInterval) {
				curInterval = mergeInterval(lastInterval, curInterval)
				res = res[:len(res)-1]
				continue
			}
			break
		}
		res = append(res, curInterval)
	}
	return res
}

func mergeInterval(a, b []int) []int {
	res := []int{0, 0}
	if len(a) != 2 || len(b) != 2 {
		return nil
	}
	// take lower
	if a[0] <= b[0] {
		res[0] = a[0]
	} else {
		res[0] = b[0]
	}

	// take higher
	if a[1] >= b[1] {
		res[1] = a[1]
	} else {
		res[1] = b[1]
	}
	return res
}

func isOverlap(a, b []int) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}
