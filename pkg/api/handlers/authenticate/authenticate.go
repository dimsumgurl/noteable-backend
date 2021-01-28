package authenticate

import (
	"github.com/dimsumgurl/noteable-backend/pkg/api/models"
	"github.com/dimsumgurl/noteable-backend/pkg/api/restapi/operations/authentication"
	"github.com/dimsumgurl/noteable-backend/pkg/domain"
	"github.com/go-openapi/runtime/middleware"
)

// AuthenticationHandler handles authentication api
type AuthenticationHandler struct {
	userService domain.UserService
}

// NewAuthenticationHandler creates a new handler for authentication
func NewAuthenticationHandler(userService domain.UserService) *AuthenticationHandler {
	return &AuthenticationHandler{userService: userService}
}

// Login will authenticate a user
func (h *AuthenticationHandler) Login(params authentication.PostLoginParams) middleware.Responder {
	token, err := h.userService.AuthenticateUser()
	if err != nil {
		// error handling
	}
	return authentication.NewPostLoginOK().WithPayload(&models.UserToken{
		Token: token,
	})
}

// Register will create a user
func (h *AuthenticationHandler) Register(params authentication.PostRegisterParams) middleware.Responder {
	err := h.userService.AddUser()
	if err != nil {
		// error handling
	}
	return authentication.NewPostRegisterOK()
}
