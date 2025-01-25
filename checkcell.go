package sudoku

func CheckCell(grid [][]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num && i != col { // updated to skip the current cell that are checking
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if grid[j][col] == num && j != row { // updated to skip the current cell that are checking
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if grid[i][j] == num && i != row && j != col { // updated to skip the current cell that are checking
				return false
			}
		}
	}
	return true
}
