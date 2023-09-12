package biz

type UserBiz struct {
}

func NewUserBiz() IUserBiz {
	return &UserBiz{}
}

func (u *UserBiz) Demo() {

}
