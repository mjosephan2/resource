package surroundedregion

func solve(board [][]byte) {
	cache := make(map[Pos]bool)
	for i, row := range board {
		for j, cell := range row {
			if cache[Pos{i, j}] {
				continue
			}
			if string(cell) == "X" {
				// solve X capture O
				capture(board, cache, [2]int{i, j})
			}
		}
	}
}

type Pos [2]int

func (p Pos) Add(change Pos) Pos {
	return Pos{p[0] + change[0], p[1] + change[1]}
}

func capture(board [][]byte, cache map[Pos]bool, pos [2]int) {
	h := len(board)
	w := len(board)

	stack := []Pos{}
	stack = append(stack, pos)
	trackingPos := []Pos{}
	directions := []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	// assume valid
	isCaptured := true

	for len(stack) > 0 {
		curPos := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cache[curPos] = true
		for _, direction := range directions {
			nextPos := curPos.Add(direction)
			if !isValid(nextPos, h, w) {
				isCaptured = false
				continue
			}
			if string(board[nextPos[0]][nextPos[1]]) != "O" {
				continue
			}
			stack = append(stack, nextPos)
			trackingPos = append(trackingPos, nextPos)
		}
	}
	if isCaptured {
		for _, pos := range trackingPos {
			board[pos[0]][pos[1]] = byte('X')
		}
	}
	return
}

func isValid(pos Pos, h, w int) bool {
	if pos[0] < 0 || pos[1] < 0 {
		return false
	}
	if pos[0] > h || pos[1] > w {
		return false
	}
	return true
}
