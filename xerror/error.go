package xerror

import (
	"errors"
)

func Is(err error, target error, others ...error) bool {
	if errors.Is(err, target) {
		return true
	}

	for _, target := range others {
		if errors.Is(err, target) {
			return true
		}
	}

	return false
}

func MessageOf(err error) string {
	var serviceErr ServiceError
	if errors.As(err, &serviceErr) {
		return serviceErr.Message
	}

	return err.Error()
}

func ServerityOf(err error, defaultSereverity Serverity) Serverity {
	var serviceErr ServiceError
	if errors.As(err, &serviceErr) {
		return serviceErr.Serverity
	}

	return defaultSereverity
}
