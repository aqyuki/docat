package display

import (
	ls "github.com/aqyuki/docat/internal/display/list"
	"github.com/aqyuki/docat/internal/display/viewport"
	"github.com/aqyuki/docat/internal/tags"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	README       = "README"
	LICENSE      = "LICENSE"
	CHANGELOG    = "CHANGELOG"
	CONTRIBUTING = "CONTRIBUTING"
	CONTRIBUTOR  = "CONTRIBUTOR"
)

var (
	// documentSelectorListItem is a content of a Document selector
	documentSelectorListItem = []list.Item{
		ls.NewModelItem(README),
		ls.NewModelItem(LICENSE),
		ls.NewModelItem(CHANGELOG),
		ls.NewModelItem(CONTRIBUTING),
		ls.NewModelItem(CONTRIBUTOR),
	}
)

// selectorList show list
func selectorList(listItem []list.Item, title string) (tea.Model, error) {
	l := list.New(listItem, ls.ItemDelegate{}, ls.DefaultWidth, ls.ListHeight)

	l.Title = title

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = ls.TitleStyle
	l.Styles.PaginationStyle = ls.PaginationStyle
	l.Styles.HelpStyle = ls.HelpStyle

	return tea.NewProgram(ls.ListModel{List: l}).Run()
}

// DocumentSelector show list to select document type
func DocumentSelector() (tags.DocumentType, error) {
	model, err := selectorList(documentSelectorListItem, "Please specify documents to be dispersed")

	if err != nil {
		return tags.NON, err
	}

	selected := ""
	switch got := model.(type) {
	case ls.ListModel:
		selected = got.Choice
	}

	switch selected {
	case README:
		return tags.README, nil
	case LICENSE:
		return tags.LICENSE, nil
	case CHANGELOG:
		return tags.CHANGELOG, nil
	case CONTRIBUTING:
		return tags.CONTRIBUTING, nil
	case CONTRIBUTOR:
		return tags.CONTRIBUTOR, nil
	default:
		return tags.NON, nil
	}
}

// FileSelector show list to select file
func FileSelector(files []string) (string, error) {
	listItems := make([]list.Item, 0)
	for _, item := range files {
		listItems = append(listItems, ls.ListItem(item))
	}

	model, err := selectorList(listItems, "Please select the file whose contents you wish to review")
	if err != nil {
		return "", err
	}

	selected := ""
	switch got := model.(type) {
	case ls.ListModel:
		selected = got.Choice
	}

	return selected, nil
}

func DocumentViewport(doc string) error {
	_, err := tea.NewProgram(
		viewport.ViewportModel{
			Content: doc,
		},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	).Run()

	if err != nil {
		return err
	}
	return nil
}
