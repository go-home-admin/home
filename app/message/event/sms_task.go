package event

import "github.com/go-home-admin/home/app/message"

type SmsTask struct {
	Phone string `json:"phone"`
}

func (receiver SmsTask) Happen() {
	message.Push(receiver)
}
