package main

import (
	"testing"
	"sandbox/barren-land"
)

func TestCreateBarrenSection(t *testing.T) {
	barrenEntry := "0 292 399 308"
	expectedBarrenSection := barrenland.BarrenSection{ 
		LowerLeft: barrenland.Location{X: 0, Y: 292}, 
		UpperRight: barrenland.Location{X:399, Y: 308},
	}
	barrenSection := barrenland.CreateBarrenSection(barrenEntry)

	if expectedBarrenSection != barrenSection {
		t.Errorf("Barren Section was incorrect, got: %d, expected: %d.", barrenSection, expectedBarrenSection)
	}
}


