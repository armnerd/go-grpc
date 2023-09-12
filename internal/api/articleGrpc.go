package api

import (
	"context"
	pb "go-grpc/internal/proto"
)

func (a *Api) GetArticleInfo(ctx context.Context, in *pb.ArticleInfoRequest) (*pb.ArticleInfoReply, error) {
	id := in.GetId()
	a.logger.Info("", id)
	info, _ := a.articleBiz.GetArticleInfo(int(id))
	a.logger.Info("", info)
	res := &pb.ArticleInfoReply{
		Id:    int32(info.ID),
		Title: info.Title,
		Intro: info.Intro,
		Cover: info.Cover,
	}
	return res, nil
}
