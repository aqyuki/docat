package view

import (
	"github.com/aqyuki/docat/internal/tags"
	ls "github.com/aqyuki/docat/internal/view/list"
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

func RunDocumentSelector() (tags.DocumentType, error) {
	l := list.New(documentSelectorListItem, ls.ItemDelegate{}, ls.DefaultWidth, ls.ListHeight)

	l.Title = "Please specify documents to be dispersed"

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = ls.TitleStyle
	l.Styles.PaginationStyle = ls.PaginationStyle
	l.Styles.HelpStyle = ls.HelpStyle

	model, err := tea.NewProgram(ls.ListModel{List: l}).Run()
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

func RunFileSelector(files []string) (string, error){
	listItems := make([]list.Item,0)
	for _, item :=range files {
		listItems = append(listItems, ls.ListItem(item))
	}
	l := list.New(listItems,ls.ItemDelegate{},ls.DefaultWidth,ls.ListHeight)

	l.Title = "Please select the file whose contents you wish to review"

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = ls.TitleStyle
	l.Styles.PaginationStyle = ls.PaginationStyle
	l.Styles.HelpStyle = ls.HelpStyle

	model,err := tea.NewProgram(ls.ListModel{List: l}).Run()
	if err != nil {
		return "",err
	}

	selected := ""
	switch got:=model.(type) {
	case ls.ListModel:
		selected = got.Choice
	}

	return selected,nil
}