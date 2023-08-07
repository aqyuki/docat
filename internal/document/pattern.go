package document

type (
	DocumentPat struct {
		Pattern string
	}
)

var (
	README = CreatePattern("README")
	LICENSE = CreatePattern("LICENSE")
	CHANGELOG = CreatePattern("CHANGELOG")
	CONTRIBUTING = CreatePattern("CONTRIBUTING")
	CONTRIBUTOR = CreatePattern("CONTRIBUTOR")
)

func CreatePattern(pattern string) *DocumentPat {
	return &DocumentPat{
		Pattern: pattern,
	}
}