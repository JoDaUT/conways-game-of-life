package main

import (
	"fmt"
	"os"
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

type Config struct {
	delay        time.Duration
	gen          int
	rows         int
	cols         int
	initialState []Point
}

func main() {

	pattern := "Random"

	config := Config{
		delay: 200 * time.Millisecond,
		gen:   40,
		cols:  30,
		rows:  30,
	}

	activeCells, err := PatternFactory(pattern, config.rows, config.cols)
	if err != nil {
		fmt.Println(fmt.Errorf("error: %w", err))
		os.Exit(1)
	}

	config.initialState = activeCells

	matrix := createWorld(config.rows, config.cols, config.initialState)
	temp := createWorld(config.rows, config.cols, []Point{})
	for i := 0; i < config.gen; i++ {
		updateState(matrix, temp, config.rows, config.cols)
		time.Sleep(config.delay)
		clearScreen()
		displayWorld(config.rows, config.cols, config.delay, i, matrix)
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
	fmt.Printf("%dx%d, delay=%dms, gen=%d\n", rows, cols, delay/time.Millisecond, iteration)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%c ", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
