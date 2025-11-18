package main

import (
	"flag"

	"github.com/raaz714/voluta/gameoflife"
	"github.com/raaz714/voluta/maze"
	"github.com/raaz714/voluta/sort"
)

func main() {
	pattern := flag.String("pattern", "maze", "Options are maze, gol, sort")
	flag.Parse()

	switch *pattern {
	case "maze":
		maze.ShowAnimatedSolution()
	case "gol":
		gameoflife.ShowAnimatedSolution()
	case "sort":
		sort.ShowAnimatedSolution()
	default:
		return
	}
}
