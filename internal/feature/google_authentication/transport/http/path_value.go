package authTransport

import (
	"fmt"
	"net/http"

	core_errors "github.com/Vlad-6894/Activity_tracking/internal/core/errors"
)

func GetStringPathValue(r *http.Request, key string) (string, error) {
	value := r.PathValue(key)
	if value == "" {
		return "", fmt.Errorf("no key='%s' in path value: %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	return value, nil
}
