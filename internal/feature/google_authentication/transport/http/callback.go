package authTransport

import (
	"net/http"

	core_logger "github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/response"
)

var (
	codeKey  = "code"
	stateKey = "state"
)

func (h *GoogleAuthHandler) GetCallback(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	code, err := GetStringPathValue(r, codeKey)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get code from path value",
		)
	}

	state, err := GetStringPathValue(r, stateKey)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get state from path value",
		)
	}

	if err := h.googleAuthService.GetTokens(ctx, code, state); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get tokens in service",
		)
	}

	responseHandler.NoContentResponse()
}
