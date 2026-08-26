package authTransport

import (
	"net/http"

	core_logger "github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/response"
)

var (
	key = "user_id"
)

type LoginResponse struct {
	Link string `json:"link"`
}

func NewLoginResponse(userLink string) LoginResponse {
	return LoginResponse{
		Link: userLink,
	}
}

func (h *GoogleAuthHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := response.NewHTTPResponseHandler(log, rw)

	userID, err := GetStringPathValue(r, key)
	if err == nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID from path value",
		)

	}

	userLink, err := h.googleAuthService.Login(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create userLink",
		)
	}

	response := NewLoginResponse(userLink)

	responseHandler.JSONResponse(response, http.StatusCreated)
}
