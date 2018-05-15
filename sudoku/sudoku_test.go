package sudoku

import (
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	testProblemJson := `{
		"Problem": [
		[1,0,3,0,0,6,0,8,0],
		[0,5,0,0,8,0,1,2,0],
		[7,0,9,1,0,3,0,5,6],
		[0,3,0,0,6,7,0,9,0],
		[5,0,7,8,0,0,0,3,0],
		[8,0,1,0,3,0,5,0,7],
		[0,4,0,0,7,8,0,1,0],
		[6,0,8,0,0,2,0,4,0],
		[0,1,2,0,4,5,0,7,8]
		],
		"Solved": false
		}`

	expectedSolutionJson := `{
		"Problem": [
		[1,2,3,4,5,6,7,8,9],
		[4,5,6,7,8,9,1,2,3],
		[7,8,9,1,2,3,4,5,6],
		[2,3,4,5,6,7,8,9,1],
		[5,6,7,8,9,1,2,3,4],
		[8,9,1,2,3,4,5,6,7],
		[3,4,5,6,7,8,9,1,2],
		[6,7,8,9,1,2,3,4,5],
		[9,1,2,3,4,5,6,7,8]
		],
		"Solved": true
		}`

	expectedSolution := convertJsonToPuzzle(expectedSolutionJson)

	solution := convertJsonToPuzzle(SolveSudoku(testProblemJson))

	if expectedSolution != solution {
		t.Errorf("Solution was incorrect, got:\n%v\nexpected:\n%v\n", solution, expectedSolution)
	}
}

//https://secure.i.telegraph.co.uk/multimedia/archive/02260/Untitled-1_2260725b.jpg
func TestWorldsHardestSudokuProblem(t *testing.T) {
	testProblemJson := `{
		"Problem": [
		[8,0,0,0,0,0,0,0,0],
		[0,0,3,6,0,0,0,0,0],
		[0,7,0,0,9,0,2,0,0],
		[0,5,0,0,0,7,0,0,0],
		[0,0,0,0,4,5,7,0,0],
		[0,0,0,1,0,0,0,3,0],
		[0,0,1,0,0,0,0,6,8],
		[0,0,8,5,0,0,0,1,0],
		[0,9,0,0,0,0,4,0,0]
		],
		"Solved": false
		}`

	expectedSolutionJson := `{
		"Problem": [
		[8,1,2,7,5,3,6,4,9],
		[9,4,3,6,8,2,1,7,5],
		[6,7,5,4,9,1,2,8,3],
		[1,5,4,2,3,7,8,9,6],
		[3,6,9,8,4,5,7,2,1],
		[2,8,7,1,6,9,5,3,4],
		[5,2,1,9,7,4,3,6,8],
		[4,3,8,5,2,6,9,1,7],
		[7,9,6,3,1,8,4,5,2]
		],
		"Solved": true
		}`

	expectedSolution := convertJsonToPuzzle(expectedSolutionJson)

	solution := convertJsonToPuzzle(SolveSudoku(testProblemJson))

	if expectedSolution != solution {
		t.Errorf("Solution was incorrect, got:\n%v\nexpected:\n%v\n", solution, expectedSolution)
	}
}
