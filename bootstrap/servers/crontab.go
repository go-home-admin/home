package servers

import (
	"github.com/robfig/cron/v3"
	"sync"
)

// Crontab @Bean
type Crontab struct {
	Cron *cron.Cron
	exit sync.Once
}

func (crontab *Crontab) Init() {
	crontab.Cron = cron.New(cron.WithSeconds())
}

func (crontab *Crontab) AddJob(spec string, cmd cron.Job) (cron.EntryID, error) {
	return crontab.Cron.AddJob(spec, cmd)
}

func (crontab *Crontab) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return crontab.Cron.AddFunc(spec, cmd)
}

func (crontab *Crontab) Run() {
	crontab.Cron.Start()
}

func (crontab *Crontab) Exit() {
	crontab.exit.Do(func() {
		crontab.Cron.Stop()
	})
}
