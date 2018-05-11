package sudoku

import (
	"strconv"
	"strings"
)

var puzzle [9][9]int

//SolveSudoku solves a 9x9 sudoku puzzle passed as a string with '_' representing blanks
func SolveSudoku(input string) string {
	puzzle = convertToPuzzle(input)

	solvePuzzle()

	return convertToString(puzzle)
}

// Convert a string into an integer array
func convertToPuzzle(input string) [9][9]int {
	var convertedPuzzle [9][9]int

	rows := strings.Split(input, "\n")
	for i := 0; i < len(rows); i++ {
		nums := strings.Split(rows[i], " ")
		for j := 0; j < len(nums); j++ {
			if nums[j] != "_" {
				num, err := strconv.Atoi(nums[j])
				if err == nil {
					convertedPuzzle[i][j] = num
				}
			}
		}
	}
	return convertedPuzzle
}

func convertToString(input [9][9]int) string {
	convertedPuzzle := ""

	for _, row := range input {
		for _, num := range row {
			convertedPuzzle += strconv.Itoa(num) + " "
		}
		convertedPuzzle += "\n"
	}

	return convertedPuzzle
}

func solvePuzzle() {

}
