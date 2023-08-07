package document

type (
	DocumentPat struct {
		Pattern string
	}
)

var (
	PatternedDocuments = map[string]*DocumentPat{
		"README":       README,
		"LICENSE":      LICENSE,
		"CHANGELOG":    CHANGELOG,
		"CONTRIBUTING": CONTRIBUTING,
		"CONTRIBUTOR":  CONTRIBUTOR,
	}

	README       = CreatePattern("README")
	LICENSE      = CreatePattern("LICENSE")
	CHANGELOG    = CreatePattern("CHANGELOG")
	CONTRIBUTING = CreatePattern("CONTRIBUTING")
	CONTRIBUTOR  = CreatePattern("CONTRIBUTOR")
)

// CreatePattern create new DocumentPat instance
func CreatePattern(pattern string) *DocumentPat {
	return &DocumentPat{
		Pattern: pattern,
	}
}
