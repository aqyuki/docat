package view

import (
	"github.com/aqyuki/docat/internal/view/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func ViewDocument(doc string) error {
	_, err := tea.NewProgram(
		viewport.ViewportModel{Content: doc},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	).Run()

	if err != nil {
		return err
	}
	return nil
}
