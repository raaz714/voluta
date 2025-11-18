package types

type (
	Coord     struct{ First, Second int }
	Neighbors map[Coord]struct{}
	AdjList   map[Coord]Neighbors
)
