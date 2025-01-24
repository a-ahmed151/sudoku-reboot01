package sudoku

func SolveGrid(grid [][]int) bool {
	/*
		this is a recursive function that will try to fill the sudoku puzzle with numbers
		and returns true if the grid is filled meaning there is a valid solution
		or false if it exhausted all posiable combinations
	*/

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 { // checks for empty spaces to fill
				for num := 1; num <= 9; num++ { // tries each number till it find one that works
					if CheckCell(grid, i, j, num) {
						grid[i][j] = num
						if SolveGrid(grid) { // recusive step to will return true if all the empty spaces are filled
							return true
						}
						grid[i][j] = 0 // reset the number to empty space if ends to deadend
					}
				}
				return false // return false if we run out of numbers to put in the empty space
			}
		}
	}
	return true // return true if all spaces are already filled
}
