package response

import "github.com/gin-gonic/gin"

type Payload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func Success(c *gin.Context, statusCode int, payload any) {
	if payload == nil {
		c.JSON(statusCode, Payload{
			Code:    statusCode,
			Message: "OK",
		})
		return
	}
	c.JSON(statusCode, payload)
}

func Error(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, Payload{
		Code:    statusCode,
		Message: err.Error(),
	})
}
