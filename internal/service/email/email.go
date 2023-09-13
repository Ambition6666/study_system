package email

import (
	"fmt"
	"net/smtp"
	"studysystem/config"
	"studysystem/internal/repository"
	tools "studysystem/pkg"

	e "github.com/jordan-wright/email"
)

// ---------------------发送验证码----------------------------------
func SendAuthCode(val ...any) any {
	v := val[0].([]any)
	to := v[0].(string)
	subject := "【study_system】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, CreateAuthCode(to))
	em := e.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = config.EmailFrom

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{to}

	// 设置主题
	em.Subject = subject

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.HTML = []byte(html)
	//fmt.Println(config.EmailAddr, config.Email, config.EmailAuth, config.EmailHost, config.EmailFrom)
	//设置服务器相关的配置
	err := em.Send(config.EmailAddr, smtp.PlainAuth("", config.Email, config.EmailAuth, config.EmailHost))
	if err != nil {
		return err
	}
	return nil
}
func CreateAuthCode(em string) string {
	code := fmt.Sprintf("%d", tools.Randnum(900000)+100000)
	repository.SetAuthCode(em, code)
	return code
}
