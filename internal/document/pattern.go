package document

import (
	"strings"

	"github.com/aqyuki/docat/internal/tags"
)

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
		"CONTRIBUTE":   CONTRIBUTE,
	}

	README       = CreatePattern("README")
	LICENSE      = CreatePattern("LICENSE")
	CHANGELOG    = CreatePattern("CHANGELOG")
	CONTRIBUTING = CreatePattern("CONTRIBUTING")
	CONTRIBUTOR  = CreatePattern("CONTRIBUTOR")
	CONTRIBUTE   = CreatePattern("CONTRIBUTE")
)

func (p *DocumentPat) Match(path string) bool {
	if isNormalArguments(path) != nil {
		return false
	}
	return compareFileName(p.Pattern, path)
}

// isNormalArguments determine if the argument can be processed normally.
func isNormalArguments(path string) error {
	if path == "" {
		return &InvalidArgument{}
	}
	return nil
}

func compareFileName(base string, name string) bool {
	upperFileName := strings.ToUpper(name)
	filename := strings.Split(upperFileName, ".")
	return strings.HasSuffix(filename[0], base)
}

// CreatePattern create new DocumentPat instance
func CreatePattern(pattern string) *DocumentPat {
	return &DocumentPat{
		Pattern: pattern,
	}
}

// FetchPatternForTag fetch pattern fot tag
func FetchPatternForTag(tag *tags.DocumentType) *DocumentPat {
	var pattern *DocumentPat
	switch *tag {
	case tags.README:
		pattern = README
	case tags.LICENSE:
		pattern = LICENSE
	case tags.CHANGELOG:
		pattern = CHANGELOG
	case tags.CONTRIBUTING:
		pattern = CONTRIBUTING
	case tags.CONTRIBUTOR:
		pattern = CONTRIBUTOR
	case tags.NON:
		pattern = nil
	}
	return pattern
}
