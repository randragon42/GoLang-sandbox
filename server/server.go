package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sandbox/sudoku"
)

func main() {
	http.HandleFunc("/sudoku", sudokuHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func sudokuHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var problem sudoku.SudokuProblem
	err := decoder.Decode(&problem)
	if err != nil {
		panic(err)

	}

	defer r.Body.Close()

	solution := sudoku.SolveSudoku(problem.Problem)
	fmt.Fprintf(w, "%v", solution)
}
