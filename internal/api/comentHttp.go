package api

import (
	"go-grpc/internal/api/response"

	"github.com/gin-gonic/gin"
)

func (api *Api) CommentDemoHttp(c *gin.Context) {
	api.commentBiz.Demo()
	response.Succuss(c, "")
}
