package login

import (
	"fmt"
	"studysystem/internal/repository"
	"studysystem/internal/service/email"
	"studysystem/internal/service/pool"
	"studysystem/models"
)

// 用户登录
func Login(em string, pwd string) (int, string) {
	u := repository.Search_user(em)
	if u.ID == 0 {
		return 401, "该用户不存在"
	}
	s := Encrypt(pwd)
	if s != u.PassWord {
		return 401, "密码错误"
	}
	token, err := GetToken(u.ID)
	if err != nil {
		return 401, "加密失败"
	}
	return 200, token
}
func Login_by_auth_code(em string, authcode string) (int, string) {
	u := repository.Search_user(em)
	if u.ID == 0 {
		return 401, "该用户不存在"
	}
	code, msg := IdentifyCode(em, authcode)
	if code != 200 {
		return code, msg
	}
	token, _ := GetToken(u.ID)
	return 200, token
}

// 用户注册
func Register(em string, pwd string) (int, string) {
	u := repository.Search_user(em)
	if u.ID != 0 {
		return 206, "用户已存在"
	}
	user := new(models.User)
	user.ID = W.GetID()
	user.Email = em
	user.PassWord = Encrypt(pwd)
	user.NickName = "默认昵称"
	user.Avatar = ""
	user.Role = 1
	repository.Create_user(user)
	return 200, "创建成功"
}

// 发送验证码
func SendAuthCode(em string) {
	pool.P.EmptyChan <- pool.NewTask(email.SendAuthCode, em)
}

// 校验验证码
func IdentifyCode(em string, authcode string) (int, string) {
	res, err := repository.GetAuthCode(em)
	if err != nil {
		fmt.Println(err)
		return 401, "验证码失效"
	}
	if res != authcode {
		return 401, "验证码错误"
	}
	return 200, "验证成功"
}
