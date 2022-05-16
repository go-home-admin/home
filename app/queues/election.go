package queues

import "github.com/go-home-admin/home/app/message"

// ElectionClose @Bean("election_close")
type ElectionClose struct {
	message.ElectionClose
}

func (e *ElectionClose) Handler() {

}
