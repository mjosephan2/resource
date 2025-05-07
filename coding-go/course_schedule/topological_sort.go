package courseschedule

import "fmt"

/*
https://leetcode.com/problems/course-schedule-ii/?envType=study-plan-v2&envId=top-interview-150
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

adj_map = map course req -> in order list of course to take before taking that course
in_degree = how many course that needs this course
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	inDegree := make([]int, numCourses)
	adj := make(map[int][]int)

	for _, pre := range prerequisites {
		adj[pre[0]] = append(adj[pre[0]], pre[1])
	}

	for _, nodes := range adj {
		for _, node := range nodes {
			inDegree[node]++
		}
	}
	fmt.Println(adj)
	fmt.Println(inDegree)
	q := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	ans := []int{}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		ans = append(ans, node)
		for _, ngbr := range adj[node] {
			inDegree[ngbr]--
			// if in degree is 0, no other course need this course, so this can be taken last
			if inDegree[ngbr] == 0 {
				q = append(q, ngbr)
			}
		}
	}

	reverse(ans)
	if len(ans) == numCourses {
		return ans
	}
	return []int{}
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
