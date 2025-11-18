package sort

import (
	"math/rand"

	"voluta/sort/types"

	"github.com/charmbracelet/x/term"
	"github.com/lucasb-eyer/go-colorful"
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

	row, col := height/2, width/2

	m := model{}
	m.G = createGrid(row, col)
	m.row = row
	m.col = col
	initiateGrid(m.G)

	return &m
}

func initiateGrid(G types.Grid) {
	for _, row := range G {
		rand.Shuffle(len(row), func(i, j int) {
			row[i], row[j] = row[j], row[i]
		})
	}
}

func createGrid(row, col int) types.Grid {
	x0y0, _ := colorful.Hex("#F27D94")
	x1y0, _ := colorful.Hex("#EDFF82")
	x0y1, _ := colorful.Hex("#643AFF")
	x1y1, _ := colorful.Hex("#14F9D5")

	x0 := make([]colorful.Color, row)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(row))
	}

	x1 := make([]colorful.Color, row)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(row))
	}

	grid := make([][]types.IndexedColor, row)
	for x := range row {
		y0 := x0[x]
		grid[x] = make([]types.IndexedColor, col)
		for y := range col {
			grid[x][y] = types.IndexedColor{X: y, Color: y0.BlendLuv(x1[x], float64(y)/float64(col)).Hex()}
		}
	}

	return grid
}
