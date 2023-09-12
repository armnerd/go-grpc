package biz

import "github.com/google/wire"

var Provider = wire.NewSet(NewArticleBiz, NewCommentBizBiz, NewUserBiz)
