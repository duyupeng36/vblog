package response

import (
	"net/http"
	"vblog/utils/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`    // 业务状态码/Http 状态码
	Message string `json:"message"` // 业务状态信息 success / failed
	Data    any    `json:"data"`
}

func NewResponse(code int, message string, data any) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SendFailed 失败
func SendFailed(c *gin.Context, err error) {
	defer c.Abort()

	switch e := err.(type) {
	case *errors.BadRequestError:
		c.JSON(http.StatusBadRequest, NewResponse(http.StatusBadRequest, "FAILED", e))
		return
	case *errors.NotFoundError:
		c.JSON(http.StatusNotFound, NewResponse(http.StatusBadRequest, "FAILED", e))
		return
	default:
		c.JSON(http.StatusInternalServerError, NewResponse(http.StatusInternalServerError, "FAILED", err))
	}

}

// SendSuccess 成功
func SendSuccess(c *gin.Context, result any) {
	c.JSON(http.StatusOK, NewResponse(http.StatusOK, "SUCCESS", result))
}
