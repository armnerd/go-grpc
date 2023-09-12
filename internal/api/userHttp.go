package api

import (
	"go-grpc/internal/api/response"

	"github.com/gin-gonic/gin"
)

func (api *Api) UserDemoHttp(c *gin.Context) {
	api.userBiz.Demo()
	response.Succuss(c, "")
}
