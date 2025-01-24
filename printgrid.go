package sudoku

import "fmt"

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Printf(" %d", row[0])
		for _, cell := range row[1:] {
			fmt.Printf(" %d", cell)
		}
		fmt.Println()
	}
}
