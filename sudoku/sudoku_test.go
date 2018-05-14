package sudoku

import "testing"

func TestSolveSudoku(t *testing.T) {
	testProblem := `1 _ 3 _ _ 6 _ 8 _
	_ 5 _ _ 8 _ 1 2 _
	7 _ 9 1 _ 3 _ 5 6
	_ 3 _ _ 6 7 _ 9 _
	5 _ 7 8 _ _ _ 3 _
	8 _ 1 _ 3 _ 5 _ 7
	_ 4 _ _ 7 8 _ 1 _
	6 _ 8 _ _ 2 _ 4 _
	_ 1 2 _ 4 5 _ 7 8`

	expectedSolution := `1 2 3 4 5 6 7 8 9
4 5 6 7 8 9 1 2 3
7 8 9 1 2 3 4 5 6
2 3 4 5 6 7 8 9 1
5 6 7 8 9 1 2 3 4
8 9 1 2 3 4 5 6 7
3 4 5 6 7 8 9 1 2
6 7 8 9 1 2 3 4 5
9 1 2 3 4 5 6 7 8
`

	solution := SolveSudoku(testProblem)

	if expectedSolution != solution {
		t.Errorf("Solution was incorrect, got:\n%v\nexpected:\n%v\n", solution, expectedSolution)
	}
}

//https://secure.i.telegraph.co.uk/multimedia/archive/02260/Untitled-1_2260725b.jpg
func TestWorldsHardestSudokuProblem(t *testing.T) {
	testProblem := `8 _ _ _ _ _ _ _ _
	_ _ 3 6 _ _ _ _ _
	_ 7 _ _ 9 _ 2 _ _
	_ 5 _ _ _ 7 _ _ _
	_ _ _ _ 4 5 7 _ _
	_ _ _ 1 _ _ _ 3 _
	_ _ 1 _ _ _ _ 6 8
	_ _ 8 5 _ _ _ 1 _
	_ 9 _ _ _ _ 4 _ _`

	expectedSolution := `8 1 2 7 5 3 6 4 9
9 4 3 6 8 2 1 7 5
6 7 5 4 9 1 2 8 3
1 5 4 2 3 7 8 9 6
3 6 9 8 4 5 7 2 1
2 8 7 1 6 9 5 3 4
5 2 1 9 7 4 3 6 8
4 3 8 5 2 6 9 1 7
7 9 6 3 1 8 4 5 2
`

	solution := SolveSudoku(testProblem)

	if expectedSolution != solution {
		t.Errorf("Solution was incorrect, got:\n%v\nexpected:\n%v\n", solution, expectedSolution)
	}
}
