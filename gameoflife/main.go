package main

import (
	"fmt"
	"time"
)

type States struct {
	ON  byte
	OFF byte
}

var STATES = States{ON: 'o', OFF: '.'}

var matrix [][]byte

type Point struct {
	X int
	Y int
}

type Patterns struct {
	Blinker         []Point
	Toad            []Point
	Beacon          []Point
	Glider          []Point
	LWSS            []Point
	Block           []Point
	Beehive         []Point
	Loaf            []Point
	GosperGliderGun []Point
}

var patterns = Patterns{
	Blinker: []Point{{1, 2}, {2, 2}, {3, 2}},
	Toad:    []Point{{2, 2}, {3, 2}, {4, 2}, {1, 3}, {2, 3}, {3, 3}},
	Beacon:  []Point{{1, 1}, {2, 1}, {1, 2}, {4, 3}, {3, 4}, {4, 4}},
	Glider:  []Point{{1, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2}},
	LWSS:    []Point{{0, 1}, {3, 1}, {4, 1}, {1, 2}, {4, 2}, {2, 3}, {3, 3}, {4, 3}, {1, 4}, {2, 4}},
	Block:   []Point{{1, 1}, {2, 1}, {1, 2}, {2, 2}},
	Beehive: []Point{{2, 1}, {3, 1}, {1, 2}, {4, 2}, {2, 3}, {3, 3}},
	Loaf:    []Point{{2, 1}, {3, 1}, {1, 2}, {4, 2}, {2, 3}, {4, 3}, {3, 4}},
	GosperGliderGun: []Point{
		{1, 5}, {1, 6}, {2, 5}, {2, 6},
		{11, 5}, {11, 6}, {11, 7},
		{12, 4}, {12, 8},
		{13, 3}, {13, 9},
		{14, 3}, {14, 9},
		{15, 6},
		{16, 4}, {16, 8},
		{17, 5}, {17, 6}, {17, 7},
		{18, 6},
		{21, 3}, {21, 4}, {21, 5},
		{22, 3}, {22, 4}, {22, 5},
		{23, 2}, {23, 6},
		{25, 1}, {25, 2}, {25, 6}, {25, 7},
		{35, 3}, {35, 4},
		{36, 3}, {36, 4},
	},
}

func main() {
	rows := 60
	cols := 60
	var delay time.Duration = 50
	gen := 1000

	activeCells := patterns.GosperGliderGun
	matrix := createWorld(rows, cols, activeCells)
	temp := createWorld(rows, cols, []Point{})
	for i := 0; i < gen; i++ {
		updateState(matrix, temp, rows, cols)
		time.Sleep(delay * time.Millisecond)
		clearScreen()
		displayWorld(rows, cols, delay, i, matrix)
	}

}

func createWorld(rows int, cols int, activeCells []Point) [][]byte {
	matrix := make([][]byte, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = STATES.OFF
		}
	}

	for _, point := range activeCells {
		matrix[point.X][point.Y] = STATES.ON
	}

	return matrix
}

func clearScreen() {
	fmt.Printf("\x1bc")
}

func isOn(matrix [][]byte, x int, y int) int {
	if matrix[x][y] == STATES.ON {
		return 1
	} else {
		return 0
	}
}

func duplicateWorld(source [][]byte, dest [][]byte) {
	rows := len(source)
	cols := len(source[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			dest[i][j] = source[i][j]
		}
	}
}

func updateState(matrix [][]byte, temp [][]byte, rows int, cols int) {
	n := 0
	duplicateWorld(matrix, temp)
	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			n = isOn(temp, i-1, j) + // up
				isOn(temp, i+1, j) + //down
				isOn(temp, i, j-1) + // left
				isOn(temp, i, j+1) + // right
				isOn(temp, i-1, j-1) + // left upper
				isOn(temp, i-1, j+1) + // right upper
				isOn(temp, i+1, j-1) + // left bottom
				isOn(temp, i+1, j+1) //right bottom

			if isOn(temp, i, j) == 1 {
				if n <= 1 || n >= 4 {
					matrix[i][j] = STATES.OFF
				}
			} else {
				if n == 3 {
					matrix[i][j] = STATES.ON
				}
			}
		}
	}
}

func displayWorld(rows int, cols int, delay time.Duration, iteration int, matrix [][]byte) {
	fmt.Printf("%dx%d, delay=%d, gen=%d\n", rows, cols, delay, iteration)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
