# Conway's Game of Life in Go

# Requirements

Navigate to https://go.dev/doc/install and choose the installer based on your operating system.

# How to build

```
cd gameoflife
go mod tidy
go build
```

# How to run

```
./gameoflife

# use help to see available options
./gameoflife -h
```

# Examples of usage

```

./gameoflife -cols 30 -rows 40 -delay 200 -pattern random -gen 200

./gameoflife -cols 30 -pattern gosperGliderGun

./gameoflife -pattern glider

```
