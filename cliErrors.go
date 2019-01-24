package main

import "github.com/pkg/errors"

const (
	ErrUndefined = ErrorType(iota)
	ErrWeb
	ErrJSON
	ErrFile
	ErrParse
)

type CliError struct {
	errorType ErrorType
	original  error
	errorText string
}

func (e CliError) Error() string {
	if e.original != nil {
		return e.original.Error()
	}
	return e.errorText
}
func (t ErrorType) New(msg string) error {
	return CliError{errorType: t, errorText: msg}
}
func (t ErrorType) Wrap(e error, msg string) error {
	newE := errors.Wrapf(e, msg)
	return CliError{errorType: t, original: newE}
}

type ErrorType uint
