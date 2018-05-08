## Barren Land Analysis - Flood Fill and Disconnected Graphs

### The Problem

You have a farm of 400m by 600m where coordinates of the field are from (0, 0) to (399, 599). A portion of the farm is barren, and all the barren land is in the form of rectangles. Due to these rectangles of barren land, the remaining area of fertile land is in no particular shape. An area of fertile land is defined as the largest area of land that is not covered by any of the rectangles of barren land. 
Read input from STDIN. Print output to STDOUT 

**Input**

You are given a set of rectangles that contain the barren land. These rectangles are defined in a string, which consists of four integers separated by single spaces, with no additional spaces in the string. The first two integers are the coordinates of the bottom left corner in the given rectangle, and the last two integers are the coordinates of the top right corner. 

**Output**

Output all the fertile land area in square meters, sorted from smallest area to greatest, separated by a space. 


### To Run
To run the Barren Land Analysis, enusure that you have GoLang installed. You can install it from https://golang.org/dl/

Once installed, clone the repo and open a terminal in the ```/sandbox``` directory.

Run ```go run main.go```
It will then prompt you to enter barren areas.

Here is an example run:
```
$ go run barren-land.go
Enter barren area: 0 292 399 307
Enter barren area:
[116800 116800]
```

In order to exit the program, simply leave an empty line and press enter (Unix).
Windows users will need to type 'done' when prompted for a barren area to signal that they have finished entering barren areas.

### To Run Test
Run ```go test``` in  the ```/barren-land``` directory to run all tests for the Barren Lands solution.

### How It Works
I made use of the [Flood Fill algorithm](https://en.wikipedia.org/wiki/Flood_fill) to solve this problem.
You start by selecting an unpainted node in a graph and recursively painting all of its neighbors the same color.
You can then count how many there are of each color and determine the size of your fertile areas.
