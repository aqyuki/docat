package options

type (
	ActivateOption func(*ScannerOption) error

	ScannerOption struct {
		IgnoreSubDirectory bool
	}
)

// CreateDefaultOption create new instance with default options
func CreateDefaultOption() *ScannerOption {
	return &ScannerOption{
		IgnoreSubDirectory: false,
	}
}

// WithIgnoreSubDirectory enable Ignore sub directory option
func WithIgnoreSubDirectory(opts *ScannerOption) {
	if opts == nil {
		return
	}
	opts.IgnoreSubDirectory = true
}
