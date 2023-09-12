package biz

import "go-grpc/internal/data/article"

type IArticleBiz interface {
	GetArticleInfo(id int) (article.Article, error)
}

type ICommentBiz interface {
	Demo()
}

type IUserBiz interface {
	Demo()
}
