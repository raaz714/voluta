package types

type (
	IndexedColor struct {
		X     int
		Y     int
		Color string
	}
	Grid [][]IndexedColor
)
