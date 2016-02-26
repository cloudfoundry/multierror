package multierror

import (
	"fmt"
	"strings"
)

type MultiError struct {
	errors []error
}

func (e MultiError) Error() string {
	aggregatedErrors := []string{errorsHeader(len(e.errors))}
	for _, err := range e.errors {
		aggregatedErrors = append(aggregatedErrors, indentLines(err))
	}
	return strings.Join(aggregatedErrors, "\n")
}

func indentLines(err error) string {
	var indentedErrors []string
	for _, line := range strings.Split(prefixErrorString(err), "\n") {
		indentedErrors = append(indentedErrors, fmt.Sprintf("    %s", line))
	}
	return fmt.Sprintf("%s", strings.Join(indentedErrors, "\n"))
}

func prefixErrorString(err error) string {
	return fmt.Sprintf("* %s", err.Error())
}

func errorsHeader(length int) string {
	var grammar string
	if length == 1 {
		grammar = "1 error"
	} else {
		grammar = fmt.Sprintf("%d errors", length)
	}

	return fmt.Sprintf("encountered %s during validation:", grammar)
}

// Add an error to the collection of errors.
// err must be non-nil
func (e *MultiError) Add(err error) {
	e.errors = append(e.errors, err)
}

// Add an error to the collection of errors with a provided prefix.
// err must be non-nil
// prefix can be empty
func (e *MultiError) AddWithPrefix(err error, prefix string) {
	errors, ok := err.(MultiError)
	if ok {
		errors.prefixAll(prefix)
		e.errors = append(e.errors, errors.errors...)
	} else {
		e.errors = append(e.errors, fmt.Errorf("%s%s", prefix, err.Error()))
	}
}

func (e *MultiError) prefixAll(prefix string) {
	for i, err := range e.errors {
		e.errors[i] = fmt.Errorf("%s%s", prefix, err.Error())
	}
}

func (e *MultiError) HasAny() bool {
	return len(e.errors) > 0
}
