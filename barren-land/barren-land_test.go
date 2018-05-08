package barrenland

import (
	"reflect"
	"testing"
)

func TestRunBarrenLandAnalysis(t *testing.T) {
	barrenSections := []string{"0 292 399 307"}
	expectedFertileSections := []int{116800, 116800}

	fertileSections := RunBarrenLandAnalysis(barrenSections)

	if !reflect.DeepEqual(expectedFertileSections, fertileSections) {
		t.Errorf("Fertile sections was incorrect, got: %d, expected %d.", fertileSections, expectedFertileSections)
	}
}

func TestRunBarrenLandAnalysis2(t *testing.T) {
	barrenSections := []string{"48 192 351 207", "48 392 351 407", "120 52 135 547", "260 52 275 547"}
	expectedFertileSections := []int{22816, 192608}

	fertileSections := RunBarrenLandAnalysis(barrenSections)

	if !reflect.DeepEqual(expectedFertileSections, fertileSections) {
		t.Errorf("Fertile sections was incorrect, got: %d, expected %d.", fertileSections, expectedFertileSections)
	}
}

func TestCreateBarrenSection(t *testing.T) {
	barrenEntry := "0 292 399 308"
	expectedBarrenSection := BarrenSection{
		LowerLeft:  Location{X: 0, Y: 292},
		UpperRight: Location{X: 399, Y: 308},
	}
	barrenSection := CreateBarrenSection(barrenEntry)

	if expectedBarrenSection != barrenSection {
		t.Errorf("Barren Section was incorrect, got: %d, expected: %d.", barrenSection, expectedBarrenSection)
	}
}
