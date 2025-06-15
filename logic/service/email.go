package service

import (
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"time"
)

var (
	ch = make(chan *gomail.Message)
)

func ListenAndSend() {
	go func() {
		defer close(ch)
		d := gomail.NewDialer("smtp.qq.com", 465, os.Getenv("username"), os.Getenv("password"))

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
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
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
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
