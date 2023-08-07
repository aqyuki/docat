package document

import (
	"path/filepath"
	"strings"
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
	}

	README       = CreatePattern("README")
	LICENSE      = CreatePattern("LICENSE")
	CHANGELOG    = CreatePattern("CHANGELOG")
	CONTRIBUTING = CreatePattern("CONTRIBUTING")
	CONTRIBUTOR  = CreatePattern("CONTRIBUTOR")
)

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

// IsREADME determine if it is a README document.
func IsREADME(path string) bool {
	err := isNormalArguments(path)
	if err != nil {
		return false
	}
	return compareFileName(README.Pattern, filepath.Base(path))
}

// IsLICENSE determine if it is a LICENSE document
func IsLICENSE(path string) bool {
	err := isNormalArguments(path)
	if err != nil {
		return false
	}
	return compareFileName(LICENSE.Pattern, filepath.Base(path))
}

// IsCHANGELOG determine if it is a CHANGELOG document
func IsCHANGELOG(path string) bool {
	err := isNormalArguments(path)
	if err != nil {
		return false
	}
	return compareFileName(CHANGELOG.Pattern, filepath.Base(path))
}

// IsCONTRIBUTING determine if it is a CONTRIBUTING document
func IsCONTRIBUTING(path string) bool {
	err := isNormalArguments(path)
	if err != nil {
		return false
	}
	return compareFileName(CONTRIBUTING.Pattern, filepath.Base(path))
}
