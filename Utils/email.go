package Utils

import (
	"gopkg.in/gomail.v2"
	"wPan/v1/Config"
)

func SendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(Config.EmailSetting.User, Config.ServerSetting.Name)) //这种方式可以添加别名，即“XX官方”
	m.SetHeader("To", mailTo...)                                                              //发送给多个用户
	m.SetHeader("Subject", subject)                                                           //设置邮件主题
	m.SetBody("text/html", body)                                                              //设置邮件正文
	d := gomail.NewDialer(
		Config.EmailSetting.Host,
		Config.EmailSetting.Port,
		Config.EmailSetting.User,
		Config.EmailSetting.Password)
	err := d.DialAndSend(m)
	return err
}
