package binutils

import (
	"fmt"
)

// Base binutils error
type Error struct {
	formattedString string
}

// Returns formatted error string as a Stringer
func (e Error) Error() string {
	return e.formattedString
}

// Returns formatted error string prefixes with package name and type as a GoStringer
func (e Error) GoString() string {
	return "binutils.Error:" + e.Error()
}

// Make new error using format strings and interface arguments
func NewError(format string, args ...interface{}) error {
	return &Error{fmt.Sprintf(format, args...)}
}
