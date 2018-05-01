package main

import (
	"fmt"
	"os"
	"bufio"
	"../sort"
	"strings"
)

type Location struct {
	x int
	y int
}

type Land struct {
	barren bool
	color int
	location Location
}

type BarrenSection struct {
	lowerLeft Location
	upperRight Location
}

type FertileSection struct {
	color int
	lands []Land
}

var AllLand [400][600]Land
var FertileSections []FertileSection
var BarrenSections []BarrenSection

func main() {
	// Initialize lists
	AllLand = [400][600]Land{}
	FertileSections = []FertileSection{}
	BarrenSections = []BarrenSection{}

	initAllLand()
	getBarrenLandsFromUser()
	FindFertileSections()

	// Sort fertile section sizes from largest to smallest
	fertileSectionSizes := []int{}
	for _, section := range FertileSections {
		fertileSectionSizes = append(fertileSectionSizes, len(section.lands))
	}
	fertileSectionSizes = sort.TopDownMergeSort(fertileSectionSizes)

	// Output list to user
	fmt.Println(fertileSectionSizes)
}

// Initialize (x, y) coordinates for all land locations
func initAllLand() {
	for i, landColumn := range AllLand {
		for j := range landColumn {
			AllLand[i][j].location = Location{i, j}
		}
	}
}

func getBarrenLandsFromUser() {
	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter barren area: ")
	input, _ := reader.ReadString('\n')

	// TODO: Due to some odd difference between the way Unix and Windows choose to give us input from the command line,
	// TODO: the following was done
	input = strings.Replace(input, "\r\n", "", -1)

	// For Unix users, hit enter on an empty barren section input line to exit. For Windows users, enter 'done'
	if strings.Compare(input, "\n") != 0 && strings.Compare(input, "done") != 0 {
		// Label barren section as barren
		newBarrenSection := CreateBarrenSection(input)
		BarrenSections = append(BarrenSections, newBarrenSection)
		MarkBarrenSection(newBarrenSection)
		getBarrenLandsFromUser()
	}


}

// Parse a string of 4 integers into a BarrenSection struct
func CreateBarrenSection(input string) (barrenSection BarrenSection){
	// Add a new Barren Section to BarrenSections
	coordinates := make([]int, 4)
	_, err := fmt.Sscan(input, &coordinates[0], &coordinates[1], &coordinates[2], &coordinates[3])
	if err != nil {
		fmt.Println(err)
		return
	}

	barrenSection = BarrenSection {
		 Location {
			coordinates[0],
			coordinates[1],
		},
		Location {
			coordinates[2],
			coordinates[3],
		},
	}

	BarrenSections = append(BarrenSections, barrenSection)
	return barrenSection
}

func MarkBarrenSection(barrenSection BarrenSection) {
	for i := barrenSection.lowerLeft.x; i <= barrenSection.upperRight.x; i++ {
		for j := barrenSection.lowerLeft.y; j <= barrenSection.upperRight.y; j++ {
			AllLand[i][j].barren = true
			AllLand[i][j].color = -1
		}
	}
}

// Find adjacent nodes that haven't been colored yet, color them, and then do the same to their neighbors
func FloodFillRecursive(land Land, section *FertileSection) {
	if land.barren { return }
	if land.color == section.color { return }
	AllLand[land.location.x][land.location.y].color = section.color
	section.lands = append(section.lands, AllLand[land.location.x][land.location.y])

	FloodFillRecursive(getNeighboringLand(land, 1, 0), section)
	FloodFillRecursive(getNeighboringLand(land, -1, 0), section)
	FloodFillRecursive(getNeighboringLand(land, 0, 1), section)
	FloodFillRecursive(getNeighboringLand(land, 0, -1), section)

}

// Find nodes that haven't been colored yet and perform a flood fill on them
func FindFertileSections(){
	color := 1
	for i := range AllLand {
		landColumn := AllLand[i]
		for j := range landColumn {
			land := AllLand[i][j]
			// Find an uncolored node
			if !land.barren && land.color == 0 {
				// Create a new fertile section
				fertileSection := FertileSection{color, []Land{}}
				// Flood Fill new fertile section starting with the uncolored node
				FloodFillRecursive(land, &fertileSection)

				FertileSections = append(FertileSections, fertileSection)
				color++	// Set up new color for next section
			}
		}
	}
}

func FloodFillQueue(land Land, section *FertileSection) {
	// Color this new section of land and add it to our fertile section
	AllLand[land.location.x][land.location.y].color = section.color
	section.lands = append(section.lands, AllLand[land.location.x][land.location.y])

	queue := make([]Land, 0)
	queue = append(queue, AllLand[land.location.x][land.location.y])

	// As long as the queue is not empty, keep flood filling
	for len(queue) != 0 {
		// Pop from queue
		x := queue[0]
		queue = queue[1:]

		// Add neighboring nodes to queue
		neighbors := []Land{
			getNeighboringLand(x, 1, 0),
			getNeighboringLand(x, -1, 0),
			getNeighboringLand(x, 0, 1),
			getNeighboringLand(x, 0, -1),
		}

		for _, neighbor := range neighbors {
			if neighbor.color == 0 {
				// Color the node, add it to our section, and append to queue
				AllLand[neighbor.location.x][neighbor.location.y].color = section.color
				section.lands = append(section.lands, AllLand[neighbor.location.x][neighbor.location.y])
				queue = append(queue, AllLand[neighbor.location.x][neighbor.location.y])
			}
		}
	}
}

// Get a neighboring land section based on offsets
// Returns an invalid land that will fail the FloodFill algorithm if the location is out of bounds
func getNeighboringLand(land Land, xOffset int, yOffset int) Land{
	invalidLand := Land{true, -1, Location{-1, -1}}

	if  land.location.x + xOffset < 0 ||
		land.location.x + xOffset > 399 ||
		land.location.y + yOffset < 0 ||
		land.location.y + yOffset > 599 {
			return invalidLand
	}

	x := land.location.x + xOffset
	y := land.location.y + yOffset
	neighbor := AllLand[x][y]

	return neighbor
}
