package email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
	"unicorn-files/pkg/logger"

	"github.com/spf13/viper"
)

type Outlook struct {
	Server   string
	Sender   string
	Password string
	Port     uint16
	Security bool
}

const (
	MailKindInvitation = "calendar"
	MailKindPlainText  = "plainText"
)

func (ol *Outlook) SendEmail(kind string, content interface{}) {
	var mailContent string
	var from string
	var to string
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         ol.Server,
	}
	auth := CustomAuth(ol.Sender, ol.Password)
	switch kind {
	case MailKindInvitation:
		mc := content.(MIMEContent)
		ctt, err := RenderTemplate(NewTemplate(kind), mc)
		//logger.Debugf("邮件内容: %s", content)
		if err != nil || content == "" {
			logger.Errorf("邮件内容模板渲染失败: %v", err)
			return
		}
		mailContent = ctt
		from = mc.From
		to = mc.To
	case MailKindPlainText:
		ptc := content.(PlainTextContent)
		ctt, err := RenderTemplate(NewTemplate(kind), ptc)
		//logger.Debugf("邮件内容: %s", content)
		if err != nil || content == "" {
			logger.Errorf("邮件内容模板渲染失败: %v", err)
			return
		}
		mailContent = ctt
		from = ptc.From
		if ptc.CC != "" {
			to = fmt.Sprintf("%v,%v", ptc.To, ptc.CC)
		} else {
			to = ptc.To
		}
	}
	c, err := smtp.Dial(ol.Server + fmt.Sprintf(":%d", ol.Port))
	if err != nil {
		logger.Errorf("SMTP 连接错误: %s", err.Error())
	} else {
		if err := c.StartTLS(tlsConf); err != nil {
			logger.Error(err)
		} else {
			if err := c.Auth(auth); err != nil {
				logger.Errorf("SMTP 认证失败: %s", err.Error())
				return
			}
			if err := c.Mail(from); err != nil {
				logger.Error(err)
				return
			}
			// 循环设置多个收件人地址
			for _, t := range strings.Split(to, ",") {
				if err := c.Rcpt(t); err != nil {
					logger.Error(err)
					return
				}
			}

			if w, err := c.Data(); err != nil {
				logger.Error(err)
				return
			} else {
				_, err = w.Write([]byte(mailContent))
				if err != nil {
					logger.Error(err)
					return
				}
				if err := w.Close(); err != nil {
					logger.Error(err)
				}
			}
			if err := c.Quit(); err != nil {
				logger.Error(err)
			}
		}
	}
}

func NewMailClient() (ol *Outlook) {
	ol = &Outlook{
		Server:   viper.GetString("email.server"),
		Sender:   viper.GetString("email.sender"),
		Password: viper.GetString("email.password"),
		Port:     uint16(viper.GetInt("email.port")),
		Security: viper.GetBool("email.security"),
	}
	return
}
