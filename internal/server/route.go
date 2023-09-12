package server

import (
	"github.com/gin-gonic/gin"
)

func route(app *gin.Engine, server *HttpServer) *gin.Engine {
	api := app.Group("/api/")
	api.POST("/article/info", server.api.GetArticleInfoHttp)
	return app
}
