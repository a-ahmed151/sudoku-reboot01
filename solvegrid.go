package sudoku

func SolveGrid(grid [][]int) bool {
	/*
		this is a recursive function that will try to fill the sudoku puzzle with numbers
		and returns true if the grid is filled meaning there is a valid solution
		or false if it exhausted all posiable combinations
	*/

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if CheckCell(grid, i, j, num) {
						grid[i][j] = num
						if SolveGrid(grid) {
							return true
						}
						grid[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}
