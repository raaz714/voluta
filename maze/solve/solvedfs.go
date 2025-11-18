package solve

import "github.com/raaz714/voluta/maze/types"

func SolveMazeDFS(
	G types.AdjList,
	visited map[types.Coord]struct{},
	solution *[]types.Coord,
	source, destin types.Coord,
) bool {
	if _, found := visited[source]; found {
		return false
	}

	visited[source] = struct{}{}
	if source == destin {
		return true
	}

	adjList, ok := G[source]
	if !ok {
		return false
	}

	for neighbor := range adjList {
		if SolveMazeDFS(G, visited, solution, neighbor, destin) {
			// solution[neighbor] = struct{}{}
			*solution = append(*solution, neighbor)
			return true
		}
	}
	return false
}
