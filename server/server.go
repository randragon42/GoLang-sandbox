package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	barrenland "sandbox/barren-land"
	"sandbox/sudoku"
)

func main() {
	http.HandleFunc("/sudoku", sudokuHandler)
	http.HandleFunc("/barren_land_analysis", barrenLandAnalysisHandler)
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

	solution := sudoku.SolveSudokuProblem(problem)
	fmt.Fprintf(w, "%v", solution)
}

func barrenLandAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var barrenSections []string
	err := decoder.Decode(&barrenSections)
	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	solution := barrenland.RunBarrenLandAnalysis(barrenSections)
	fmt.Fprintf(w, "%v", solution)
}
