package deadlock

/**

 * Deadlock is a situation in which two or more processes are unable to proceed because each is waiting for the other to release a resource.
 * In Go, deadlocks can occur when goroutines are waiting on each other to release locks or channels.
 * This can happen with mutexes, channels, or any synchronization primitives.

 * To avoid deadlocks, ensure that all goroutines acquire locks in a consistent order, use timeouts, or leverage Go's built-in concurrency patterns.

 Questions for the interview:
 Find if this resource has a deadlock
 Input:
 [
   ["A", "B"],
   ["B", "C"],
   ["C", "A"],
   ["C", "D"],
 ]
**/

func FindDeadlock(graph [][]string) bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var dfs func(string) bool
	dfs = func(node string) bool {
		if !visited[node] {
			visited[node] = true
			recStack[node] = true

			for _, neighbor := range getNeighbors(graph, node) {
				if !visited[neighbor] && dfs(neighbor) {
					return true
				} else if recStack[neighbor] {
					return true
				}
			}
		}
		recStack[node] = false
		return false
	}

	for _, edge := range graph {
		if dfs(edge[0]) {
			return true
		}
	}

	return false
}

// Consider using a recursive function method to traverse the graph
func getNeighbors(graph [][]string, node string) []string {
	neighbors := []string{}
	for _, edge := range graph {
		if edge[0] == node {
			neighbors = append(neighbors, edge[1])
		}
	}
	return neighbors
}