package sort

import (
	"github.com/raaz714/voluta/sort/types"

	"github.com/charmbracelet/lipgloss"
)

func printGrid(G types.Grid) string {
	res := ""
	for _, x := range G {
		for _, y := range x {
			s := lipgloss.NewStyle().SetString("  ").Background(lipgloss.Color(y.Color))
			res += s.String()
		}
		res += "\n"
	}
	return res
}

// bubbleSortPass performs one pass of the bubble sort algorithm
func bubbleSortPass(G types.Grid) bool {
	swapped := false // Flag to track if any swaps occurred in this pass

	for i := 0; i < len(G)-1; i++ {
		// Compare adjacent elements
		if G[i][0].X > G[i+1][0].X {
			// Swap if they are in the wrong order
			G[i], G[i+1] = G[i+1], G[i]
			swapped = true
		}
	}
	return swapped // Return true if a swap occurred, indicating the array is not yet sorted
}

func bubbleSortPassRow(G types.Grid) bool {
	swapped := false // Flag to track if any swaps occurred in this pass

	for _, row := range G {
		for i := 0; i < len(row)-1; i++ {
			// Compare adjacent elements
			if row[i].Y > row[i+1].Y {
				// Swap if they are in the wrong order
				row[i], row[i+1] = row[i+1], row[i]
				swapped = true
			}
		}
	}
	return swapped // Return true if a swap occurred, indicating the array is not yet sorted
}

func insertionSortPass(G types.Grid) bool {
	swapped := false
	for _, row := range G {
		indexToInsert := 1
		for row[indexToInsert].Y >= row[indexToInsert-1].Y {
			indexToInsert++
			if indexToInsert >= len(row) {
				break
			}
		}

		if indexToInsert <= 0 || indexToInsert >= len(row) {
			// Nothing to insert or invalid index
			continue
		}

		key := row[indexToInsert] // The value to be inserted
		j := indexToInsert - 1    // Start comparing with the element just before the key

		// Move elements of the sorted subarray that are greater than key
		// one position ahead of their current position
		for j >= 0 && row[j].Y > key.Y {
			row[j+1] = row[j]
			j--
		}

		// Place the key in its correct position
		row[j+1] = key
		swapped = true
	}
	return swapped
}

func updateGrid(G types.Grid) bool {
	flag1 := bubbleSortPass(G)
	flag2 := /*bubbleSortPassRow(G) */ insertionSortPass(G)
	return flag1 || flag2
}
