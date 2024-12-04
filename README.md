# advent-of-code-go-template
A very simple template for completing [Advent Of Code](https://adventofcode.com) challenges in Golang.

## Usage

The template uses an interface to the functions `A()` and `B()` for each day's challenge. 
These can be imported into the `main.go` file as show. 

A `Makefile` is provided to facilitate the generation and cleanup of new days and an empty input.txt file.

Once generated, the day can be imported into the [main.go](main.go) file.

[Learn more about creating a repository from a template](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template).

## Generating a new Day's Challenge

Generate a Day 2challenge

```shell
make gen d=2
```

This creates a new folder for the provided day, as well as a go file and an empty input.txt file. It is assumed you will
copy the contents of the challenge's input file manually.

Once created, update the imports in the `main.go` file accordingly:

```golang
import (
    ...
    "github.com/mikezm/advent-of-code-2024/day2"
)
...
...
var challengeMap = challenges{
	1: day1.Challenge{},
	2: day2.Challenge{},
}
...
```

This removes the day's folder and all it's contents. 

## Executing a day's challenge

```shell
go run main.go 1 a
```

```shell
go run main.go 1 B
```
 
