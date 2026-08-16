package main

import "fmt"

type Notification interface {
	Send(message string)
}

type EmailNotification struct{}

func (e EmailNotification) Send(message string) {
	fmt.Println("Send Email Notification: ", message)
}

type SMSNotification struct{}

func (e SMSNotification) Send(message string) {
	fmt.Println("Send Sms Notification:", message)
}

func notify(n Notification, message string) {
	n.Send(message)
}
func main() {
	email := EmailNotification{}
	sms := SMSNotification{}

	notify(email, "12345")
	notify(sms, "12345")

}

var _ Notification = (*EmailNotification)(nil)
var _ Notification = (*SMSNotification)(nil)
