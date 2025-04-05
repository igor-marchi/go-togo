package todo

var ErrInvalidTitle = Error("title is required")

type Error string

func (e Error) Error() string { return string(e) }
