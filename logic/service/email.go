package service

import (
	"context"
	log "github.com/faiz/llm-code-review/common/logger"
	"github.com/faiz/llm-code-review/config"
	"gopkg.in/gomail.v2"
	"time"
)

var (
	ch = make(chan *gomail.Message)
)

func ListenAndSend(ctx context.Context) {
	log.New(ctx).Info("starting email service")
	go func() {
		defer close(ch)
		// TODO username 和 password 从环境变量中读取
		d := gomail.NewDialer("smtp.qq.com", 465, config.Email.Username, config.Email.Password)

		var s gomail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					log.New(ctx).Info("email service stopped")
					return
				}
				if !open {
					log.New(ctx).Debug("opening email connection")
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := gomail.Send(s, m); err != nil {
					log.New(ctx).Error("failed to send email", "message", m, "error", err)
				}
				log.New(ctx).Info("sent email", "message", m)
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				log.New(ctx).Info("closing email connection")
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
