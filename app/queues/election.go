package queues

import (
	"github.com/go-home-admin/home/app/message"
	"github.com/go-home-admin/home/bootstrap/services/app"
)

// ElectionClose @Bean("election_close")
type ElectionClose struct {
	message.ElectionClose
}

type run interface {
	Run()
}

func (e *ElectionClose) Handler() {
	// 立即唤醒
	app.GetBean("election").(run).Run()
}
