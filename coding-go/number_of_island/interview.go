package island

func numIslandsNoEdge(grid [][]byte) int {
	// similar to numIsland but this time we do not consider the case where the land touches the edge of the box
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				if isIslandNoEdge(grid, i, j) {
					count++
				}

			}
		}
	}

	return count
}

func isIslandNoEdge(grid [][]byte, i, j int) (res bool) {
	res = true
	q := [][2]int{}
	q = append(q, [2]int{i, j})

	boundaryI := len(grid)
	boundaryJ := len(grid[0])

	offsets := [4][2]int{
		{-1, 0}, // above
		{1, 0},  // below
		{0, 1},  // right
		{0, -1}, // left
	}
	for len(q) > 0 {
		curr := q[len(q)-1]
		q = q[:len(q)-1]

		// set it to explored
		grid[curr[0]][curr[1]] = '2'

		for _, offset := range offsets {
			newI := curr[0] + offset[0]
			newJ := curr[1] + offset[1]

			// check for boundaries
			if newI >= 0 && newI < boundaryI && newJ >= 0 && newJ < boundaryJ {
				// check if it is a land that has not been explored
				if grid[newI][newJ] != '1' {
					continue
				}
				q = append(q, [2]int{newI, newJ})
			} else {
				// it is not an island
				res = false
			}
		}
	}
	return
}
