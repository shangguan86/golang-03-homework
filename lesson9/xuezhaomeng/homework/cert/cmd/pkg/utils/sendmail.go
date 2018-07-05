package utils

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

const (
	HOST        = "mail.xxxx.com"
	SERVER_ADDR = "mail.xxxx.com:25"
	USER        = "xxxxx" //发送邮件的邮箱
	PASSWORD    = "xxxxx" //发送邮件邮箱的密码
)

type Email struct {
	to      string "to"
	subject string "subject"
	msg     string "msg"
}

func NewEmail(to, subject, msg string) *Email {
	return &Email{to: to, subject: subject, msg: msg}
}

func SendEmail(email *Email) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	//fmt.Println(auth)
	sendTo := strings.Split(email.to, ";")
	done := make(chan error, 1024)

	go func() {
		defer close(done)
		for _, v := range sendTo {

			str := strings.Replace("From: "+USER+"~To: "+v+"~Subject: "+email.subject+"~~", "~", "\r\n", -1) + email.msg

			err := SmtpSendMail(
				SERVER_ADDR,
				auth,
				USER,
				[]string{v},
				[]byte(str),
			)
			fmt.Println(err)
			done <- err
		}
	}()

	for i := 0; i < len(sendTo); i++ {
		<-done
	}

	return nil
}

func SmtpSendMail(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
	//fmt.Print(addr, a, from, to, msg)
	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	//  if err = c.hello(); err != nil {
	//      return err
	//  }
	if ok, _ := c.Extension("STARTTLS"); ok {
		host, _, _ := net.SplitHostPort(addr)
		config := &tls.Config{InsecureSkipVerify: true, ServerName: host}
		//      if testHookStartTLS != nil {
		//          testHookStartTLS(config)
		//      }
		if err = c.StartTLS(config); err != nil {
			return err
		}
	}
	//  if a != nil && c.ext != nil {
	//      if _, ok := c.ext["AUTH"]; ok {
	if a != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(a); err != nil {
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
