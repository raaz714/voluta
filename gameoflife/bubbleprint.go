package gameoflife

import (
	"fmt"
	"time"

	"voluta/gameoflife/types"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

// Model represents the application state
type model struct {
	row int
	col int
	G   types.Grid
}

func delayTick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// tickMsg is a custom message type for timer ticks
type tickMsg time.Time

// Init starts the ticker command.
func (m model) Init() tea.Cmd {
	return delayTick()
}

// Update handles incoming messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		// Handle key presses for exiting.
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "r" {
			initiateGrid(m.G, m.row, m.col)
		}
	case tickMsg:
		updateGrid(m.G, m.row, m.col)
		return m, delayTick()
	}

	return m, nil
}

// View renders the UI
func (m model) View() string {
	// s := "Press 's' to start/stop, 'q' to quit.\n\n"

	// Get the terminal size
	physicalWidth, physicalHeight, err := term.GetSize(0)
	if err != nil {
		// Handle error, e.g., if not in a terminal, set a default width
		physicalWidth = 80
	}
	gridString := printGrid(m.G, m.row, m.col)
	gridStringColor := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")). // Purple color
		Width(physicalWidth - 4).
		Height(physicalHeight - 4).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center).
		Render(gridString)
	return gridStringColor
}

func ShowAnimatedSolution() {
	p := tea.NewProgram(createNewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}
