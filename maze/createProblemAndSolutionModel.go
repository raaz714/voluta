package maze

import (
	"math/rand"

	"voluta/maze/populate"
	"voluta/maze/solve"
	"voluta/maze/types"

	"github.com/charmbracelet/x/term"
)

func createNewModel() *model {
	if !term.IsTerminal(0) {
		println("not in a term")
		return nil
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return nil
	}

	G := make(types.AdjList)
	visited := make(map[types.Coord]struct{})
	row, col := height/3-2, width/4-2

	populate.PopulateMazeDFS(types.Coord{First: 5, Second: 5}, G, visited, row, col)

	source := types.Coord{First: rand.Intn(row), Second: rand.Intn(col)}
	destin := types.Coord{First: rand.Intn(row), Second: rand.Intn(col)}

	visited = make(map[types.Coord]struct{})
	solution := []types.Coord{}

	solve.SolveMazeDFS(G, visited, &solution, destin, source)
	solutionMap := make(map[types.Coord]struct{})

	return &model{row, col, source, destin, G, solution, 0, solutionMap}
}
