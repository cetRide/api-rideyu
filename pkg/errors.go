package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Error struct {
	err            error  // original error
	errorMessage   string // Contains error message for error that can be displayed to user
	httpStatusCode int
	logMessages    []string // Contains the "stack" of error messages
}

// Error interface
func (e *Error) Error() string {
	return fmt.Sprintf("[ERROR] %s", strings.Join(e.logMessages, "; "))
}

func NewError(err error, format string, a ...interface{}) *Error {

	appError, ok := err.(*Error)
	if !ok {
		return NewErrorWithCodeAndMessage(
			err,
			http.StatusInternalServerError,
			"",
			format,
			a...,
		)
	}

	appError.addLogMessage(err, false, format, a...)

	return appError
}

func NewErrorWithCodeAndMessage(
	err error,
	httpStatus int,
	message string,
	format string,
	a ...interface{},
) *Error {

	appError, ok := err.(*Error)
	if !ok {
		appError = &Error{
			err: err,
		}

		appError.httpStatusCode = http.StatusInternalServerError
	}

	appError.httpStatusCode = httpStatus

	if message != "" {
		appError.errorMessage = message
	}

	appError.addLogMessage(err, true, format, a...)

	return appError
}

func (e *Error) HttpStatus() int {

	if e.httpStatusCode != 0 {
		return e.httpStatusCode
	}

	return http.StatusBadRequest
}

func IsErrNoRows(err error) bool {

	appError, ok := err.(*Error)

	if !ok {
		return err == sql.ErrNoRows
	} else {
		return appError.err == sql.ErrNoRows
	}
}

func (e *Error) JsonResponse() map[string]string {

	if e.errorMessage == "" {
		e.errorMessage = "Failed to perform request. Please try again."
	}

	return map[string]string{
		"error_message": e.errorMessage,
	}
}

func (e *Error) LogErrorMessages() {
	log.Println(e.Error())
}

func (e *Error) addLogMessage(err error, includeErr bool, format string, a ...interface{}) {

	suffix := ""
	if includeErr {
		suffix = fmt.Sprintf(" because err=[%v]", err)
	}

	e.logMessages = append([]string{fmt.Sprintf("%s%s", fmt.Sprintf(format, a...), suffix)}, e.logMessages...)
}
