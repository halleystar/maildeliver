package service

import (
	"fmt"
	"log"

	"potato/utils"
	"time"

	"gopkg.in/gomail.v2"
)

type Email struct {
	msgQueue chan *gomail.Message
}

type Message struct {
	From    string `json:"from"`
	TO      string `json:"to" binding:"required"`
	Cc      string `json:"cc"`
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
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
	fmt.Println("end.....")
	//	msg.SetHeader("To", rawMsg.TO)
	//	msg.SetAddressHeader("Cc", rawMsg.Cc, "")
	//	msg.SetHeader("Subject", rawMsg.Subject)
	//	msg.SetBody("text/html", rawMsg.Body)
	msg.SetHeader("To", "1247920356@qq.com")
	msg.SetHeader("Subject", "subject")
	msg.SetBody("text/html", "hahahha")
	return msg
}

func (email Email) sendMessage() {
	go func() {
		d := gomail.NewDialer(utils.Cfg.EmailHost, utils.Cfg.Port, utils.Cfg.Username, utils.Cfg.Password)
		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-email.msgQueue:
				if !ok {
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						panic(err)
					}
					open = false
				}
			}
		}
	}()
}
