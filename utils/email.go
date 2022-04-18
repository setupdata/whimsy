package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"whimsy/global"
)

// SendEmailWithTLS SSL/TLS 发送邮件服务
func SendEmailWithTLS(to string, code string) {
	from := global.PIC_CONFIG.EMAIL.From
	subj := "奇妙博客注册邮箱验证码" // 主题
	body := "邮箱验证码为: " + code
	// 设置 headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subj

	// 设置 message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// 连接 SMTP服务器
	hostname := global.PIC_CONFIG.EMAIL.Host

	host, _, _ := net.SplitHostPort(hostname)

	auth := smtp.PlainAuth("", from, global.PIC_CONFIG.EMAIL.SmtpPassword, host)

	// TLS 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	// 建立tcp连接
	conn, err := tls.Dial("tcp", hostname, tlsconfig)
	if err != nil {
		global.PIC_LOG.Error("邮件服务tls连接错误", err)
		return
	}
	// 创建邮件客户端
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		global.PIC_LOG.Error("邮件服务新建客户端错误", err)
		return
	}

	// 认证
	if err = c.Auth(auth); err != nil {
		global.PIC_LOG.Error("邮件服务认证错误", err)
		return
	}

	// To && From
	if err = c.Mail(from); err != nil {
		global.PIC_LOG.Error("邮件服务启动邮件事务错误", err)
		return
	}

	if err = c.Rcpt(to); err != nil {
		global.PIC_LOG.Error("邮件服务Rcpt调用错误", err)
		return
	}

	// 数据
	w, err := c.Data()
	if err != nil {
		global.PIC_LOG.Error("邮件服务编写数据错误", err)
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		global.PIC_LOG.Error("邮件服务写入数据错误", err)
		return
	}

	err = w.Close()
	if err != nil {
		global.PIC_LOG.Error("邮件服务关闭writer错误", err)
		return
	}

	err = c.Quit()
	if err != nil {
		global.PIC_LOG.Error("邮件服务关闭服务器连接错误", err)
		return
	}
}
