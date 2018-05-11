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
9 1 2 3 4 5 6 7 8`

	solution := SolveSudoku(testProblem)

	if expectedSolution != solution {
		t.Errorf("Solution was incorrect, got:\n%v\nexpected:\n%v\n", solution, expectedSolution)
	}
}
