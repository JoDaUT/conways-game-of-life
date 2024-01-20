# Conway's Game of Life in Go

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

![random-pattern](https://github.com/JoDaUT/conways-game-of-life/assets/47344349/8443e8bc-92ea-4bad-8daa-93d82449f448)

`./gameoflife -pattern gosperGliderGun -gen 200`

![gosperGliderGun-pattern](https://github.com/JoDaUT/conways-game-of-life/assets/47344349/8a19a25a-aa39-410a-b743-da10a25cdf48)

`./gameoflife -pattern lwss -rows 8

![lwss-pattern](https://github.com/JoDaUT/conways-game-of-life/assets/47344349/fc76ab66-da5c-4992-8f7a-b8063d7fd775)
