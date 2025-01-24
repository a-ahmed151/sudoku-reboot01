package main

import (
	"fmt"
	"os"
	"sudoku"
)

func main() {
	input := os.Args
	grid, valid := sudoku.ValidInput(input)
	if !valid {
		fmt.Println("ERROR")
		return
	}
	if sudoku.SolveGrid(grid) {
		sudoku.PrintGrid(grid)
	} else {
		fmt.Println("ERROR")
		return
	}
}
