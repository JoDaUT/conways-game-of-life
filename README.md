# Conway's Game of Life in Go

## Requirements

Download and install Go on your machine: https://go.dev/doc/install

## How to build

Go to the project folder and execute the following commands:

```
cd gameoflife
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

```

./gameoflife -cols 30 -rows 40 -delay 200 -pattern random -gen 200

./gameoflife -cols 30 -pattern gosperGliderGun

./gameoflife -pattern glider

```
