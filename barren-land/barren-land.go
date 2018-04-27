package main

import (
	"fmt"
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
	AllLand = [400][600]Land{}
	FertileSections = []FertileSection{}
	BarrenSections = []BarrenSection{}

	input := "0 292 399 307"

	newBarrenSection := CreateBarrenSection(input)

	fmt.Println("Barren Sections: ", BarrenSections)
	MarkBarrenSection(newBarrenSection)

	color := 1
	for i, landColumn := range AllLand {
		for j, land := range landColumn {
			land.location = Location{i, j}
			if !land.barren && land.color == 0 {
				fertileSection := FertileSection{color, []Land{}}
				FloodFill(land, fertileSection)

				FertileSections = append(FertileSections, fertileSection)
				color++
			}
		}
	}

	fmt.Println("Number of fertile sections", len(FertileSections))
}

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

// Recursive FloodFill causes a stack overflow :(
func FloodFillRecursive(land Land, section FertileSection) {
	if land.barren { return }
	if land.color == section.color { return }
	land.color = section.color
	section.lands = append(section.lands, land)

	FloodFill(getNeighboringLand(land, 1, 0), section)
	FloodFill(getNeighboringLand(land, -1, 0), section)
	FloodFill(getNeighboringLand(land, 0, 1), section)
	FloodFill(getNeighboringLand(land, 0, -1), section)

}

func FloodFill(land Land, section FertileSection) {
	if land.barren { return }
	if land.color == section.color { return }
	land.color = section.color
	section.lands = append(section.lands, land)

	queue := make([]Land, 0)
	queue = append(queue, land)

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
				AllLand[neighbor.location.x][neighbor.location.y].color = section.color
				section.lands = append(section.lands, AllLand[neighbor.location.x][neighbor.location.y])
				queue = append(queue, AllLand[neighbor.location.x][neighbor.location.y])
			}
		}

		// Remove - kick us out if
		if len(queue) > 240000 {
			fmt.Println("Your algorithm isn't right...")
			panic(fmt.Sprintf("Queue is larger than maximum possible. Fix your algorithm"))
		}
	}

}

func getNeighboringLand(land Land, xOffset int, yOffset int) Land{
	invalidLand := Land{true, -1, Location{-1, -1}}

	if  land.location.x + xOffset < 0 ||
		land.location.x + xOffset > 399 ||
		land.location.y + yOffset < 0 ||
		land.location.y + yOffset > 599 {
			return invalidLand
	}

	//fmt.Println(land.location.x + xOffset, land.location.y + yOffset)
	return AllLand[land.location.x + xOffset][land.location.y + yOffset]
}
