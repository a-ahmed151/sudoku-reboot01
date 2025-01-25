package sudoku

func ValidInput(input []string) ([][]int, bool) {
	/*
	   this function will take the argument given and try to validate it
	   and return a grid and true if the input is valid
	   or returns a empty grid and false if the input is invalid
	*/

	if len(input) != 10 {
		return nil, false
	}
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
	}
	for i, r := range input[1:] {
		if len(r) != 9 {
			return nil, false
		}

		for j, char := range r {
			if !(char >= '1' && char <= '9' || char == '.') {
				return nil, false
			} else if char == '.' {
				grid[i][j] = 0
			} else {
				grid[i][j] = int(char - '0')
			}
		}
	}

	// this checks if the inputed grid is correct with no repated numbers in row, column and boxes
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := grid[row][col]
			if !CheckCell(grid, row, col, num) && num != 0 {
				return nil, false
			}
		}
	}

	return grid, true
}
