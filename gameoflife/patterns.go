package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

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
	Random          []Point
}

var patternsMap = map[string][]Point{
	"Blinker":         patterns.Blinker,
	"Toad":            patterns.Toad,
	"Beacon":          patterns.Beacon,
	"Glider":          patterns.Glider,
	"LWSS":            patterns.LWSS,
	"Block":           patterns.Block,
	"Beehive":         patterns.Beehive,
	"Loaf":            patterns.Loaf,
	"GosperGliderGun": patterns.GosperGliderGun,
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

func generatePseudoRandomPattern(size int, rows int, cols int) []Point {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	points := []Point{}

	counter := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if counter <= size && random.Float64() > 0.5 {
				points = append(points, Point{X: i, Y: j})
			}
			counter++
		}
	}
	return points
}

func PatternFactory(pattern string, rows int, cols int) ([]Point, error) {
	var patternKeys = []string{}
	for k := range patternsMap {
		patternKeys = append(patternKeys, k)
	}
	patternKeys = append(patternKeys, "Random")

	if pattern == "Random" {
		return generatePseudoRandomPattern(rows*cols, rows, cols), nil
	}
	if found, ok := patternsMap[pattern]; !ok {

		return nil, fmt.Errorf("invalid pattern '%s': allowed patterns'%s'", pattern, strings.Join(patternKeys, ","))
	} else {
		return found, nil
	}
}
