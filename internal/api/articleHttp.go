package api

import (
	"go-grpc/internal/api/dto"

	"go-grpc/internal/api/response"

	"github.com/gin-gonic/gin"
)

func (a *Api) GetArticleInfoHttp(c *gin.Context) {
	var request dto.ArticleInfoRequest
	c.ShouldBind(&request)
	if request.Id == 0 {
		response.Fail(c, response.ParamsLost)
		return
	}
	a.logger.Info("", request.Id)
	info, _ := a.articleBiz.GetArticleInfo(request.Id)
	a.logger.Info("", info)
	res := &dto.ArticleInfoReply{
		Id:    info.ID,
		Title: info.Title,
		Intro: info.Intro,
		Cover: info.Cover,
	}
	response.Succuss(c, res)
}
