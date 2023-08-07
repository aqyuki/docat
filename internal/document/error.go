package document

type (
	NotFindError    struct{}
	InvalidArgument struct{}
)

func (e *NotFindError) Error() string {
	return "no file matching the pattern was found"
}

func NewNotFindError() error {
	return &NotFindError{}
}

func (e *InvalidArgument) Error() string {
	return "invalid argument"
}

func NewInvalidArgument() error {
	return &InvalidArgument{}
}
