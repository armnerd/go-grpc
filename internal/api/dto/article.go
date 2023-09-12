package dto

type ArticleInfoRequest struct {
	Id int `form:"id" json:"id" `
}

type ArticleInfoReply struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Intro  string `json:"intro"`
	Cover  string `json:"cover"`
	Author string `json:"author"`
}
