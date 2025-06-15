package service

import (
	"gopkg.in/gomail.v2"
	"os"
	"testing"
	"time"
)

func TestEmail(t *testing.T) {
	ListenAndSend()
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("username"))
	m.SetHeader("Subject", "TestEmail")
	m.SetHeader("To", "")
	m.SetBody("text/html", "Hello World!")
	ch <- m
	<-time.After(time.Second)
}
