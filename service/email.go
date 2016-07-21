package service

import (
	"fmt"
	"log"

	"maildeliver/utils"
	"time"

	"gopkg.in/gomail.v2"
)

type Email struct {
	msgQueue chan *gomail.Message
}

type Message struct {
	TO          []string `json:"to" binding:"required"`
	Cc          []string `json:"cc"`
	Subject     string   `json:"subject" binding:"required"`
	Body        string   `json:"body" binding:"required"`
	ContentType int      `json:"content_type" binding:"required"`
	IsSplitTo   bool     `json:"is_split_to" binding:"required"`
}

func initEmail() *Email {
	email := &Email{}
	email.msgQueue = make(chan *gomail.Message)
	email.sendMessage()
	return email
}

func (e Email) Send(rawMsg Message) {
	msg := e.initMessage(rawMsg)
	e.msgQueue <- msg
}

func (e Email) initMessage(rawMsg Message) *gomail.Message {
	msg := gomail.NewMessage()
	if rawMsg.From == "" {
		msg.SetHeader("From", utils.Cfg.FromEmail)
	} else {
		msg.SetHeader("From", rawMsg.From)
	}
	msg.SetHeader("To", rawMsg.TO)
	msg.SetAddressHeader("Cc", rawMsg.Cc, "")
	msg.SetHeader("Subject", rawMsg.Subject)
	msg.SetBody("text/html", rawMsg.Body)
	//	msg.SetHeader("To", "1247920356@qq.com")
	//	msg.SetHeader("Subject", "subject")
	//	msg.SetBody("text/html", "hahahha")
	return msg
}

/*
发送邮箱核心服务, 支持并行发送，该方法负责发送每封邮件内容不想同的邮件
*/
func (email Email) sendMessage() {
	dialer := gomail.NewDialer(utils.Cfg.EmailHost, utils.Cfg.Port, utils.Cfg.Username, utils.Cfg.Password)
	var sendCloser gomail.SendCloser
	ticket := utils.NewTicket()
	open := false
	for {
		select {
		case msg, ok := <-email.msgQueue:
			if !ok {
				return
			}
			if !open {
				if sendCloser, err := dialer.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			ticket.Done()
			go func() {
				defer ticket.Add()
				if err := gomail.Send(sendCloser, msg); err != nil {
					log.Print(err)
				}
			}()
		case <-time.After(30 * time.Second):
			if open {
				if err := sendCloser.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}
