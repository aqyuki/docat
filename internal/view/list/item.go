package list

type (
	ListItem string
)

func (m ListItem) FilterValue() string {
	return ""
}

func NewModelItem(label string) ListItem {
	return ListItem(label)
}
