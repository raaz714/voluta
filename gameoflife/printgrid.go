package gameoflife

import "voluta/gameoflife/types"

func printGrid(G types.Grid, row, col int) string {
	res := ""
	for r := range row {
		for c := range col {
			if G[r][c] == 0 {
				res += "  "
			} else {
				res += "⬜"
			}
		}
		res += "\n"
	}
	return res
}

func updateGrid(G types.Grid, row, col int) {
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	for i := range row {
		for j := range col {
			live := 0
			for _, dir := range directions {
				x, y := (i+dir[0]+row)%row, (j+dir[1]+row)%row
				if x >= 0 && x < row && y >= 0 && y < col && (G[x][y] == 1 || G[x][y] == 3) {
					live++
				}
			}

			if G[i][j] == 1 && (live < 2 || live > 3) {
				G[i][j] = 3
			} else if G[i][j] == 0 && live == 3 {
				G[i][j] = 2
			}
		}
	}

	for i := range row {
		for j := range col {
			if G[i][j] == 2 {
				G[i][j] = 1
			}
			if G[i][j] == 3 {
				G[i][j] = 0
			}
		}
	}
}
