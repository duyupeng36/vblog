package middleware

import (
	"vblog/apps/user"
	"vblog/ioc"
	"vblog/logger"
	"vblog/utils/errors"
	"vblog/utils/response"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	// 中间件：认证中间件
	// 1. 获取 token
	logger.Logger().Info().Msg("auth ....")

	tkString, err := c.Cookie(user.AUTOH_COOKIE_NAME)
	if err != nil {
		response.SendFailed(c, errors.NewUnauthorizedError("Not Login, please login first"))
		return
	}

	// 2. 调用 user 的 CheckToken 校验 token 的合法性
	controller := ioc.Container().GetController(user.AppName).(user.Service)

	tk, err := controller.CheckToken(c.Request.Context(), user.NewCheckTokenRequest(tkString))
	if err != nil {
		response.SendFailed(c, errors.NewUnauthorizedError("Not Login, please login first"))
		return
	}
	// 3. 把 User 对象放入 gin.Context 中。后续的接口可能需要从 上下文中获取用户
	c.Set(user.REQUEST_CTX_TOKEN_NAME, tk)
	logger.Logger().Info().Msg("authed sucessfull ....")
}
