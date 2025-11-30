package delivery

import (
	"errors"
	"manager/internal/usecase"
	"manager/shared/response"
	"net/http"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Login struct {
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h *handleImpl) HandleLogin(c *gin.Context) {
	var loginInfo Login
	if err := c.ShouldBind(&loginInfo); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(loginInfo); err != nil {
		response.Error(c, http.StatusBadRequest, err)
		return
	}

	result, err := h.LoginUseCase.Execute(c.Request.Context(), usecase.LoginInput{
		Phone:    loginInfo.Phone,
		Password: loginInfo.Password,
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, errors.New("StatusNotFound"))
			return
		}
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	response.Success(c, http.StatusOK, result)
}
