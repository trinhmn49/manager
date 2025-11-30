package delivery

import (
	"manager/internal/delivery/model/req"
	"manager/internal/usecase"
	"manager/shared/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handleImpl) HandleRegister(c *gin.Context) {
	var user req.CreateUser
	if err := c.ShouldBind(&user); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	if err := h.RegisterUseCase.Execute(
		c.Request.Context(),
		usecase.RegisterInput{
			DisplayName: user.DisplayName,
			Avatar:      user.Avatar,
			Phone:       user.Phone,
			Password:    user.Password,
			Email:       user.Email,
		},
	); err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusCreated, nil)
}
