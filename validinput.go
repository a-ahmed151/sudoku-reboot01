package sudoku

func ValidInput(input []string) ([][]int, bool) {
	/*
	   this function will take the argument given and try to validate it
	   and return a grid and true if the input is valid
	   or returns a empty grid and false if the input is invalid
	*/
	if len(input) != 9 {
		return nil, false
	}
	grid := make([][]int, 9)
	for i := 0; i < 9; i++ {
		grid[i] = make([]int, 9)
	}
	for i, r := range input {
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
	return grid, true
}
