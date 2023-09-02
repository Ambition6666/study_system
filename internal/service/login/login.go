package login

import (
	"studysystem/internal/repository"
	"studysystem/models"
)

// 用户登录
func Login(email string, pwd string) (int, string) {
	u := repository.Search_user(email)
	if u.ID == 0 {
		return 401, "该用户不存在"
	}
	s := Encrypt(pwd)
	if s != u.PassWord {
		return 401, "密码错误"
	}
	token, err := GetToken(Msk, u.ID)
	if err != nil {
		return 401, "加密失败"
	}
	return 200, token
}

// 用户注册
func Register(email string, pwd string) (int, string) {
	u := repository.Search_user(email)
	if u.ID != 0 {
		return 206, "用户已存在"
	}
	user := new(models.User)
	user.ID = W.GetID()
	user.Email = email
	user.PassWord = Encrypt(pwd)
	user.NickName = "默认昵称"
	user.Avatar = ""
	repository.Create_user(user)
	return 200, "创建成功"
}
