package apperrors

import "errors"

type Kind uint8

const (
	InvalidInput Kind = iota
	Unauthorized
	Domain
	NotFound
	Conflict
	Internal
)

type Error struct {
	kind    Kind
	message string
}

func New(
	kind Kind,
	message string,
) *Error {
	return &Error{
		kind,
		message,
	}
}

func (e *Error) Kind() Kind {
	return e.kind
}

func (e *Error) Error() string {
	return e.message
}

func KindOf(e error) Kind {
	var err *Error

	if errors.As(e, &err) {
		return err.kind
	}

	return Internal
}
