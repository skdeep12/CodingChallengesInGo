package UrlShortner

import "fmt"

type Error interface {
	Error() string
	Code() string
	Message() string
	EmbeddedError() string
}

type ErrorImpl struct {
	code          string
	message       string
	embeddedError Error
}

func NewError(code string, message string) Error {
	return &ErrorImpl{code: code, message: message}
}

func WrapError(code string, message string, err Error) Error {
	return &ErrorImpl{code: code, message: message, embeddedError: err}
}

func (e *ErrorImpl) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("[%s,%s]\n%s", e.Code(), e.Message(), e.EmbeddedError())
}

func (e *ErrorImpl) Code() string {
	return e.code
}
func (e *ErrorImpl) Message() string {
	return e.message
}
func (e *ErrorImpl) EmbeddedError() string {
	if e.embeddedError != nil {
		return e.embeddedError.Error()
	}
	return ""
}
