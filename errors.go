package binutils

import (
	"fmt"
)

// Error type implements error interface and used in package routines.
type Error struct {
	wrapped     error
	description string
}

// Error returns formatted error string. Implements error interface.
func (e Error) Error() string {
	if e.wrapped != nil {
		return e.description + ": " + e.wrapped.Error()
	}

	return e.description
}

// GoString returns formatted error string prefixes with package name and type. Implements GoStringer.
func (e Error) GoString() string {
	return "binutils.Error:" + e.Error()
}

// NewError is a default error constructor.
// Takes format and arguments to make error description using fmt.Sprintf() formatting rules.
func NewError(format string, args ...interface{}) error {
	return &Error{description: fmt.Sprintf(format, args...)}
}

// WrapError creates new error, wrapping underlying error with formatted string.
func WrapError(wrap error, format string, args ...interface{}) *Error {
	return &Error{description: fmt.Sprintf(format, args...), wrapped: wrap}
}
