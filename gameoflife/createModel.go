package gameoflife

import (
	"math/rand"

	"voluta/gameoflife/types"

	"github.com/charmbracelet/x/term"
)

func initiateGrid(G types.Grid, row, col int) {
	for i := range row {
		for j := range col {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				G[i][j] = 1
			} else {
				G[i][j] = rand.Intn(2)
			}
		}
	}
}

func createNewModel() *model {
	if !term.IsTerminal(0) {
		println("not in a term")
		return nil
	}
	width, height, err := term.GetSize(0)
	if err != nil {
		return nil
	}

	row, col := height*5/6, width/2-8

	m := model{}
	// Initialize the outer slice (rows)
	m.G = make([][]int, row)

	// Initialize each inner slice (columns)
	for i := range m.G {
		m.G[i] = make([]int, col)
	}
	m.row = row
	m.col = col
	initiateGrid(m.G, m.row, m.col)

	return &m
}
