package sudoku

import (
	"strconv"
	"strings"
)

type SudokuProblem struct {
	Problem string
}

type Spot struct {
	row int
	col int
}

var puzzle [9][9]int

//SolveSudoku solves a 9x9 sudoku puzzle passed as a string with '_' representing blanks
func SolveSudoku(input string) string {
	puzzle = convertToPuzzle(input)

	solvePuzzleBruteForce()

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

// Convert an integer array representing a sudoku puzzle into a multiline string
func convertToString(input [9][9]int) string {
	convertedPuzzle := ""

	for _, row := range input {
		for j, num := range row {
			convertedPuzzle += strconv.Itoa(num)
			if j != len(row)-1 {
				convertedPuzzle += " "
			}
		}
		convertedPuzzle += "\n"
	}

	return convertedPuzzle
}

func solvePuzzleBruteForce() {
	unsolvedSpots := []Spot{}
	// Loop through the puzzle looking for unsolved spaces and store in slice
	for i, row := range puzzle {
		for j, num := range row {
			if num == 0 {
				unsolvedSpots = append(unsolvedSpots, Spot{i, j})
			}
		}
	}

	solveSpot(unsolvedSpots, 0)
}

// Recursively brute force solutions for the puzzle until we find one or exhaust all possible options
func solveSpot(unsolvedSpots []Spot, index int) bool {
	spot := unsolvedSpots[index]
	// Find all possible values for the index passed in
	possibleSolutions := findPossibleSolutions(unsolvedSpots[index])
	// If we are at the last index and there is only one value -> end condition -> return true
	if index == len(unsolvedSpots)-1 && len(possibleSolutions) == 1 {
		puzzle[spot.row][spot.col] = possibleSolutions[0]
		return true
	}

	// Loop through all of those values and try them
	for _, possibleSolution := range possibleSolutions {
		puzzle[spot.row][spot.col] = possibleSolution
		// Call solveSpot(array, index + 1)
		// if it returns true, a solution was found
		if solveSpot(unsolvedSpots, index+1) {
			return true
		}
	}

	// Reset the spot to 0 (spot still unsolved)
	puzzle[spot.row][spot.col] = 0
	return false

}

func findPossibleSolutions(spot Spot) []int {
	// Start with a slice of values 1 through 9
	possibleSolutions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Check row
	for _, num := range puzzle[spot.row] {
		if num != 0 {
			possibleSolutions = removeFromPossibleSolutions(num, possibleSolutions)
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if puzzle[i][spot.col] != 0 {
			possibleSolutions = removeFromPossibleSolutions(puzzle[i][spot.col], possibleSolutions)
		}
	}

	// Check box
	x := spot.row - (spot.row % 3)
	y := spot.col - (spot.col % 3)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if puzzle[i][j] != 0 {
				possibleSolutions = removeFromPossibleSolutions(puzzle[i][j], possibleSolutions)
			}
		}
	}

	return possibleSolutions
}

func removeFromPossibleSolutions(numberToRemove int, possibleSolutions []int) []int {
	for i, num := range possibleSolutions {
		if num == numberToRemove {
			return append(possibleSolutions[:i], possibleSolutions[i+1:]...)
		}
	}
	return possibleSolutions
}
