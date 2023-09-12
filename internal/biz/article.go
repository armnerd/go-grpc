package biz

import (
	"go-grpc/internal/data/article"
)

type ArticleBiz struct {
	ArticleDao article.IArticleDao
}

func NewArticleBiz(article article.IArticleDao) IArticleBiz {
	return &ArticleBiz{
		ArticleDao: article,
	}
}

func (a *ArticleBiz) GetArticleInfo(id int) (article.Article, error) {
	return a.ArticleDao.GetOne(id)
}
