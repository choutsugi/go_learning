package routes

import (
	"net/http"
	"web_app/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup() (r *gin.Engine) {
	r = gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("./version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.App.Version)
	})
	return
}
