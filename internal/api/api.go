package api

import (
	"go-grpc/internal/biz"

	"go-grpc/internal/base"
	pb "go-grpc/internal/proto"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewApi, biz.Provider)

type Api struct {
	pb.UnimplementedBlogServer
	articleBiz biz.IArticleBiz
	commentBiz biz.ICommentBiz
	userBiz    biz.IUserBiz
	logger     base.Logger
}

func NewApi(articleBiz biz.IArticleBiz, commentBiz biz.ICommentBiz, userBiz biz.IUserBiz, logger base.Logger) *Api {
	return &Api{
		articleBiz: articleBiz,
		commentBiz: commentBiz,
		userBiz:    userBiz,
		logger:     logger,
	}
}
