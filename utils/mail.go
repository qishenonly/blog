package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"path"

	"github.com/qishenonly/blog/global"

	"gopkg.in/gomail.v2"
)

// 邮件发送
func Send(to, subject, body string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", global.Config.Email.User)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// 设置 163 邮箱 SMTP 服务器配置
	dialer := gomail.NewDialer(global.Config.Email.Host, global.Config.Email.Port,
		global.Config.Email.User, global.Config.Email.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// 发送邮件
	if err := dialer.DialAndSend(m); err != nil {
		global.Logger.Errorf("邮件发送失败: %s", err)
	}

	return err

}

type VerifyTokenData struct {
	Token string
	Email string
}

// SendWithTemplate 解析并渲染邮件模板
func SendWithTemplate(tplName string, data interface{}, to, subject string) error {
	tplFile := "utils/view/" + tplName + ".gohtml"
	t, err := template.New(path.Base(tplFile)).ParseFiles(tplFile)
	if err != nil {
		return err
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, data); err != nil {
		return err
	}
	return Send(to, subject, tpl.String())
}
