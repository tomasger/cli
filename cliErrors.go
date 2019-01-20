package main

import "github.com/pkg/errors"

const (
	ErrUndefined = ErrorType(iota)
	ErrNetwork
	ErrHTTP
	ErrJSON
	ErrFile
)
type cliError struct {
	errorType ErrorType
	original error
	errorText string
}
func (e cliError) Error() string {
	s := ""
	switch e.errorType {
	case ErrUndefined:
		s = "(ErrUndefined)"
	case ErrNetwork:
		s = "(ErrNetwork)"
	case ErrHTTP:
		s = "(ErrHTTP)"
	case ErrJSON:
		s = "(ErrJSON)"
	}
	s = e.original.Error() + " " + s
	return e.original.Error()
}
func (t ErrorType) New(msg string) error {
	return cliError{errorType:t, errorText:msg}
}
func (t ErrorType) Wrap(e error, msg string) error {
	newE := errors.Wrapf(e, msg)
	return cliError{errorType: t, original:newE}
}
type ErrorType uint