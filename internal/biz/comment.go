package biz

type CommentBiz struct {
}

func NewCommentBizBiz() ICommentBiz {
	return &CommentBiz{}
}

func (c *CommentBiz) Demo() {

}
