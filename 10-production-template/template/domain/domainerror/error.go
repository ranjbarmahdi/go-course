package domainerror

import "errors"

type Kind uint8

const (
	InvalidInput Kind = iota
	NotFound
	Conflict
	Domain
	Internal
)

type Error struct {
	kind    Kind
	message string
}

func New(kind Kind, message string) *Error {
	return &Error{kind: kind, message: message}
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
