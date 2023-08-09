package tags

type (
	DocumentType int
)

const (
	NON          DocumentType = -1
	README       DocumentType = 0
	LICENSE      DocumentType = 1
	CHANGELOG    DocumentType = 2
	CONTRIBUTING DocumentType = 3
	CONTRIBUTOR  DocumentType = 4
	CONTRIBUTE   DocumentType = 5
)
