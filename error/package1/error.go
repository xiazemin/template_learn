package package1

import "errors"

// Create a named type for our new error type.
type errorString string

// Implement the error interface
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType error = errors.New("EOF")
