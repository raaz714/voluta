package maze

import (
	"fmt"
	"strings"

	"voluta/maze/types"

	"github.com/charmbracelet/lipgloss"
)

// Unicode characters for single-line box drawing
const (
	cornerTopLeft     = "┌"
	cornerTopRight    = "┐"
	cornerBottomLeft  = "└"
	cornerBottomRight = "┘"
	tJunctionDown     = "┬"
	tJunctionUp       = "┴"
	tJunctionRight    = "├"
	tJunctionLeft     = "┤"
	cross             = "┼"
	horizontal        = "─"
	vertical          = "│"
	space             = " "
)

var (
	sourceStyle = lipgloss.NewStyle().Background(lipgloss.Color("#0000FF"))
	destinStyle = lipgloss.NewStyle().Background(lipgloss.Color("#04B575"))
	pathStyle   = lipgloss.NewStyle().Background(lipgloss.Color("#F1E0B6"))
)

func DrawGrid(rows, cols int, source, destin types.Coord, G types.AdjList, solution map[types.Coord]struct{}) {
	for r := range rows {
		// Print top border of a cell row or horizontal separator
		for c := range cols {
			fmt.Print("+")
			edges, ok := G[types.Coord{First: r - 1, Second: c}]
			if r != 0 && ok {
				if _, found := edges[types.Coord{First: r, Second: c}]; found {
					fmt.Print("   ")
					continue
				}
			}
			fmt.Print("───")
		}
		fmt.Println("+")

		// Print content of a cell row
		for c := range cols {
			edge, ok := G[types.Coord{First: r, Second: c - 1}]
			if c != 0 && ok {
				if _, found := edge[types.Coord{First: r, Second: c}]; found {
					fmt.Print(" ")
				} else {
					fmt.Print("│")
				}
			} else {
				fmt.Print("│")
			}
			_, found := solution[types.Coord{First: r, Second: c}]
			switch {
			case source == (types.Coord{First: r, Second: c}):
				fmt.Print(" s ")
			case destin == (types.Coord{First: r, Second: c}):
				fmt.Print(" d ")
			case found:
				fmt.Print(" * ")
			default:
				fmt.Print("   ")
			}
		}
		fmt.Println("│")
	}

	// Print bottom border of the last row
	for range cols {
		fmt.Print("+───")
	}
	fmt.Println("+")
}

func GridString(rows, cols int, source, destin types.Coord, G types.AdjList, solution map[types.Coord]struct{}) string {
	var res string
	for r := range rows {
		// Print top border of a cell row or horizontal separator
		for c := range cols {
			res += "+"
			edges, ok := G[types.Coord{First: r - 1, Second: c}]
			if r != 0 && ok {
				if _, found := edges[types.Coord{First: r, Second: c}]; found {
					res += "   "
					continue
				}
			}
			res += "───"
		}
		res += "+\n"

		// Print content of a cell row
		for c := range cols {
			edge, ok := G[types.Coord{First: r, Second: c - 1}]
			if c != 0 && ok {
				if _, found := edge[types.Coord{First: r, Second: c}]; found {
					res += " "
				} else {
					res += "│"
				}
			} else {
				res += "│"
			}
			_, found := solution[types.Coord{First: r, Second: c}]
			switch {
			case source == (types.Coord{First: r, Second: c}):
				res += "⭕ "
			case destin == (types.Coord{First: r, Second: c}):
				res += "🚪 "
			case found:
				res += " * "
			default:
				res += "   "
			}
		}
		res += "│\n"
	}

	// Print bottom border of the last row
	for range cols {
		res += "+───"
	}
	res += "+\n"

	return res
}

func GridStringUnicode(rows, cols int, source, destin types.Coord, G types.AdjList, solution map[types.Coord]struct{}) string {
	var res string

	// Horizontal line segment (cell width)
	hSegment := strings.Repeat(horizontal, 3) // Use 3 horizontal chars for cell width
	// Cell content space (cell width)
	// cSegmentEmpty := strings.Repeat(space, 3)

	// Build the top border
	topBorder := cornerTopLeft + strings.Repeat(hSegment+tJunctionDown, cols-1) + hSegment + cornerTopRight + "\n"
	res += topBorder

	for r := range rows {
		// Build the cell row
		// cellRow1 := vertical + strings.Repeat(cSegmentEmpty+vertical, cols) + "\n"
		// cellRow1 := cellRow(G, solution, source, destin, r, cols, false)
		cellRow2 := cellRow(G, solution, source, destin, r, cols, true)
		res += cellRow2 + cellRow2

		// Build the separator row (between cells)
		if r < rows-1 {
			// separatorRow := tJunctionRight + strings.Repeat(hSegment+cross, cols-1) + hSegment + tJunctionLeft + "\n"
			separatorRow := cellBorder(G, r, cols)
			res += separatorRow
		}
	}

	// Build the bottom border
	bottomBorder := cornerBottomLeft + strings.Repeat(hSegment+tJunctionUp, cols-1) + hSegment + cornerBottomRight + "\n"
	res += bottomBorder
	// fmt.Println(res)
	return res
}

func edgeExists(G types.AdjList, r1, c1, r2, c2 int) bool {
	edges, ok := G[types.Coord{First: r1, Second: c1}]
	if ok {
		if _, found := edges[types.Coord{First: r2, Second: c2}]; found {
			return true
		}
	}
	return false
}

func getConnector(G types.AdjList, r, c int) string {
	aboveOk := !edgeExists(G, r, c, r, c+1)
	belowOk := !edgeExists(G, r+1, c, r+1, c+1)
	leftOk := !edgeExists(G, r, c, r+1, c)
	rightOk := !edgeExists(G, r, c+1, r+1, c+1)

	if aboveOk && belowOk && leftOk && rightOk {
		return cross
	}

	if aboveOk {
		if belowOk {
			if leftOk {
				return tJunctionLeft
			} else if rightOk {
				return tJunctionRight
			} else {
				return "│"
			}
		} else {
			if leftOk {
				if rightOk {
					return tJunctionUp
				} else {
					return cornerBottomRight
				}
			} else if rightOk {
				return cornerBottomLeft
			}
		}
	}

	if belowOk {
		if leftOk {
			if rightOk {
				return tJunctionDown
			} else {
				return cornerTopRight
			}
		} else {
			if rightOk {
				return cornerTopLeft
			}
		}
	}
	if leftOk {
		if rightOk {
			return "‒"
		}
	}
	return " "
}

func cellBorder(G types.AdjList, r, cols int) string {
	hSegment := strings.Repeat(horizontal, 3) // Use 3 horizontal chars for cell width
	cellBorderRow := tJunctionRight
	for c := range cols {
		if edgeExists(G, r, c, r+1, c) {
			cellBorderRow += "   "
			if c != cols-1 {
				cellBorderRow += getConnector(G, r, c)
			}
			continue
		}
		cellBorderRow += hSegment
		if c != cols-1 {
			cellBorderRow += getConnector(G, r, c) // cross
		}
	}
	cellBorderRow += tJunctionLeft + "\n"

	return cellBorderRow
}

func cellRow(G types.AdjList, solution map[types.Coord]struct{}, source, destin types.Coord, r, cols int, cellContent bool) string {
	cSegmentEmpty := strings.Repeat(space, 3)
	cSegmentPath := pathStyle.Render("   ") // " * "
	cSegmentStart := sourceStyle.Render("   ")
	cSegmentDest := destinStyle.Render("   ")
	cellRow := ""

	for c := range cols {
		if c != 0 && edgeExists(G, r, c, r, c-1) {
			cellRow += " "
		} else {
			cellRow += vertical
		}
		if cellContent {
			_, found := solution[types.Coord{First: r, Second: c}]
			switch {
			case source == (types.Coord{First: r, Second: c}):
				cellRow += cSegmentStart
			case destin == (types.Coord{First: r, Second: c}):
				cellRow += cSegmentDest
			case found:
				cellRow += cSegmentPath
			default:
				cellRow += cSegmentEmpty
			}
		} else {
			cellRow += cSegmentEmpty
		}
	}
	return cellRow + vertical + "\n"
}
