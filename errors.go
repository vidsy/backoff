package backoff

import "strings"

type (
	// Errors slice of errors.
	Errors []error
)

// Error implements the error interface.
func (e Errors) Error() string {
	var errors []string
	for _, err := range e {
		errors = append(
			errors,
			err.Error(),
		)
	}

	return strings.Join(errors, ", ")
}
