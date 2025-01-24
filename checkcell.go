package sudoku

func CheckCell(grid [][]int, row int, col int, num int) bool {
    for i := 0; i < 9; i++ {
        if grid[row][i] == num || grid[i][col] == num {
            return false
        }
    }

    startRow := (row / 3) * 3
    startCol := (col / 3) * 3
    for i := startRow; i < startRow+3; i++ {
        for j := startCol; j < startCol+3; j++ {
            if grid[i][j] == num {
                return false
            }
        }
    }
    return true
}