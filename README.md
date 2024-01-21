# Conway's Game of Life in Go

It is a famous cellular automaton proposed by the mathematician John Horton Conway in 1970. It requires no more interaction than the initial state of the game.



The rules are really simple:

- If a cell has less than two live neighbors, it dies of loneliness.

- If a cell has two or three live neighbors, it lives on to the next generation.

- If a cell has more than three live neighbors, it dies of overpopulation.



With these pretty simple rules a lot of funny and interesting patterns have been found so far, e.g. blocks, beacon, glider, glider guns and space ships, just to mention the most famous ones.

## Requirements

Download and install Go on your machine: https://go.dev/doc/install

## How to build

Go to the project folder and execute the following commands:

```
go mod tidy
go build
```

## How to run

```
# linux and macos
./gameoflife

# windows
./gameoflife.exe

# use help to see available options
./gameoflife -h
```

## Examples of usage

`./gameoflife -cols 50 -rows 40 -delay 200 -pattern random -gen 200`

![random-pattern](https://github.com/JoDaUT/gameoflife/assets/47344349/f2ed6f41-bb74-48f4-ba50-002a91f11048)


`./gameoflife -pattern gosperGliderGun -gen 200`

![gosperGliderGun-pattern](https://github.com/JoDaUT/gameoflife/assets/47344349/c151cebf-d9d3-416a-8c4c-bf1070a1cf99)


`./gameoflife -pattern lwss -rows 8`

![lwss-pattern](https://github.com/JoDaUT/gameoflife/assets/47344349/d19221d2-8390-4dda-b646-25d0e728484d)



