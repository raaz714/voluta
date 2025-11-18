package populate

import (
	"math/rand"

	"github.com/raaz714/voluta/maze/types"
)

var directions = []int{0, -1, 0, 1, 0}

func PopulateMazeDFS(curr types.Coord, G types.AdjList, visited map[types.Coord]struct{}, row, col int) {
	visited[curr] = struct{}{}
	unvisited_neighbors := make(types.Neighbors)

	for i := range 4 {
		x := curr.First + directions[i]
		y := curr.Second + directions[i+1]

		if x < 0 || x == row || y < 0 || y == col {
			continue
		}

		unvisited_neighbors[types.Coord{First: x, Second: y}] = struct{}{}
	}

	for {
		if len(unvisited_neighbors) == 0 {
			return
		}

		// Select random neighbor
		randIdx := rand.Intn(len(unvisited_neighbors))
		var selected types.Coord
		i := 0
		for k := range unvisited_neighbors {
			if i == randIdx {
				selected = k
				break
			}
			i++
		}

		if _, found := visited[selected]; found {
			delete(unvisited_neighbors, selected)
			continue
		}

		if _, ok := G[curr]; !ok {
			G[curr] = make(types.Neighbors)
		}
		G[curr][selected] = struct{}{}

		if _, ok := G[selected]; !ok {
			G[selected] = make(types.Neighbors)
		}
		G[selected][curr] = struct{}{}

		delete(unvisited_neighbors, selected)

		PopulateMazeDFS(selected, G, visited, row, col)
	}
}
