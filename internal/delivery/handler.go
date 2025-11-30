package delivery

import (
	// "manager/shared/middleware"
	"manager/internal/usecase"
	"manager/pkg/server"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	server.RouteRegister
	HandleRegister(c *gin.Context)
	HandleLogin(c *gin.Context)
	HandleGetUserProfile(c *gin.Context)
}

type handleImpl struct {
	LoginUseCase    usecase.LoginUseCase
	RegisterUseCase usecase.RegisterUseCase
}

// TODO
func (h *handleImpl) RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/user/login", h.HandleLogin)
	r.POST("/user/register", h.HandleRegister)
}

func NewCustomerHandler(LoginUseCase usecase.LoginUseCase, RegisterUseCase usecase.RegisterUseCase) Handler {
	return &handleImpl{
		LoginUseCase:    LoginUseCase,
		RegisterUseCase: RegisterUseCase,
	}
}
