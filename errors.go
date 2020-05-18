package binutils

import (
	"fmt"
)

// Error type implements error interface and used in package routines.
type Error struct {
	formattedString string
}

// Error returns formatted error string. Implements error interface.
func (e Error) Error() string {
	return e.formattedString
}

// GoString returns formatted error string prefixes with package name and type. Implements GoStringer
func (e Error) GoString() string {
	return "binutils.Error:" + e.Error()
}

// NewError is a default error constructor.
// Takes format and arguments to make error description using fmt.Sprintf() formatting rules.
func NewError(format string, args ...interface{}) error {
	return &Error{fmt.Sprintf(format, args...)}
}
