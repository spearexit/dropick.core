package apis

import (
	"github.com/gin-gonic/gin"
)

//	 Ping godoc
//		@Summary	ping to check server status
//		@Schemes
//		@Description	do ping
//		@Tags			/health
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	JSONResult
//		@Router			/health/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, JSONResult{
		Code:    200,
		Message: "pong",
		Data:    nil,
	})
}
