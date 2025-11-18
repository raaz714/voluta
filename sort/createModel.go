package sort

import (
	"fmt"
	"math/rand"

	"github.com/raaz714/voluta/sort/types"

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

	row := min(height*5/6, width/2-10)
	col := row

	m := model{}
	m.G = createGrid(row, col)
	m.row = row
	m.col = col
	initiateGrid(m.G)

	return &m
}

func initiateGrid(G types.Grid) {
	row := len(G)
	col := len(G[0])
	// x0y0, _ := colorful.Hex("#F27D94")
	// x1y0, _ := colorful.Hex("#EDFF82")
	// x0y1, _ := colorful.Hex("#643AFF")
	// x1y1, _ := colorful.Hex("#14F9D5")
	x0y0, _ := colorful.Hex(generateRandomHexColor())
	x1y0, _ := colorful.Hex(generateRandomHexColor())
	x0y1, _ := colorful.Hex(generateRandomHexColor())
	x1y1, _ := colorful.Hex(generateRandomHexColor())

	x0 := make([]colorful.Color, row)
	for i := range x0 {
		x0[i] = x0y0.BlendLuv(x0y1, float64(i)/float64(row))
	}

	x1 := make([]colorful.Color, row)
	for i := range x1 {
		x1[i] = x1y0.BlendLuv(x1y1, float64(i)/float64(row))
	}
	for x := range row {
		y0 := x0[x]
		for y := range col {
			G[x][y] = types.IndexedColor{X: x, Y: y, Color: y0.BlendLuv(x1[x], float64(y)/float64(col)).Hex()}
		}
	}
	rand.Shuffle(len(G), func(i, j int) {
		G[i], G[j] = G[j], G[i]
	})

	for _, row := range G {
		rand.Shuffle(len(row), func(i, j int) {
			row[i], row[j] = row[j], row[i]
		})
	}
}

func generateRandomHexColor() string {
	// Generate three random integers for R, G, and B components (0-255)
	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	// Format the integers as a hexadecimal string
	hexColor := fmt.Sprintf("#%02x%02x%02x", r, g, b)
	return hexColor
}

func createGrid(row, col int) types.Grid {
	grid := make([][]types.IndexedColor, row)
	for x := range row {
		grid[x] = make([]types.IndexedColor, col)
	}
	initiateGrid(grid)

	return grid
}
