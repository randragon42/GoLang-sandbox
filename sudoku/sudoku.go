package sudoku

import (
	"encoding/json"
	"fmt"
)

type SudokuProblem struct {
	Problem [9][9]int
	Solved  bool
}

type Spot struct {
	row int
	col int
}

var puzzle SudokuProblem

//SolveSudoku solves a 9x9 sudoku puzzle passed as a string with '_' representing blanks
func SolveSudoku(input string) string {
	puzzle = ConvertJsonToPuzzle(input)

	puzzle.Solved = solvePuzzleBruteForce()

	return ConvertPuzzleToJson(puzzle)
}

func ConvertJsonToPuzzle(input string) SudokuProblem {
	sudokuProblem := SudokuProblem{}
	err := json.Unmarshal([]byte(input), &sudokuProblem)
	if err != nil {
		fmt.Println(err.Error())
	}
	return sudokuProblem
}

func ConvertPuzzleToJson(input SudokuProblem) string {
	problemJson, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(problemJson)
}

func solvePuzzleBruteForce() bool {
	unsolvedSpots := []Spot{}
	// Loop through the puzzle looking for unsolved spaces and store in slice
	for i, row := range puzzle.Problem {
		for j, num := range row {
			if num == 0 {
				unsolvedSpots = append(unsolvedSpots, Spot{i, j})
			}
		}
	}

	return solveSpot(unsolvedSpots, 0)
}

// Recursively brute force solutions for the puzzle until we find one or exhaust all possible options
func solveSpot(unsolvedSpots []Spot, index int) bool {
	spot := unsolvedSpots[index]
	// Find all possible values for the index passed in
	possibleSolutions := findPossibleSolutions(unsolvedSpots[index])
	// If we are at the last index and there is only one value -> end condition -> return true
	if index == len(unsolvedSpots)-1 && len(possibleSolutions) == 1 {
		puzzle.Problem[spot.row][spot.col] = possibleSolutions[0]
		return true
	}

	// Loop through all of those values and try them
	for _, possibleSolution := range possibleSolutions {
		puzzle.Problem[spot.row][spot.col] = possibleSolution
		// Call solveSpot(array, index + 1)
		// if it returns true, a solution was found
		if solveSpot(unsolvedSpots, index+1) {
			return true
		}
	}

	// Reset the spot to 0 (spot still unsolved)
	puzzle.Problem[spot.row][spot.col] = 0
	return false

}

func findPossibleSolutions(spot Spot) []int {
	// Start with a slice of values 1 through 9
	possibleSolutions := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Check row
	for _, num := range puzzle.Problem[spot.row] {
		if num != 0 {
			possibleSolutions = removeFromPossibleSolutions(num, possibleSolutions)
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if puzzle.Problem[i][spot.col] != 0 {
			possibleSolutions = removeFromPossibleSolutions(puzzle.Problem[i][spot.col], possibleSolutions)
		}
	}

	// Check box
	x := spot.row - (spot.row % 3)
	y := spot.col - (spot.col % 3)
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			if puzzle.Problem[i][j] != 0 {
				possibleSolutions = removeFromPossibleSolutions(puzzle.Problem[i][j], possibleSolutions)
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
