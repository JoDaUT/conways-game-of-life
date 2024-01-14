package main

import (
	"fmt"
	"time"
)

var matrix [][]byte

func main() {
	rows := 10
	cols := 15
	var delay time.Duration = 1000
	gen := 20

	matrix = make([][]byte, rows)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]byte, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = '.'
		}
	}

	for i := 0; i < gen; i++ {
		time.Sleep(delay * time.Millisecond)
		fmt.Printf("\x1bc")

		displayWorld(rows, cols, delay, i, matrix)
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
