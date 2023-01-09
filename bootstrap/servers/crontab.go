package servers

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type Cron func(job func())

// Crontab @Bean
type Crontab struct {
	Cron *cron.Cron
	exit sync.Once

	middlewares []Cron
}

func (crontab *Crontab) Init() {
	crontab.Cron = cron.New(cron.WithSeconds())
}

func (crontab *Crontab) AddJob(spec string, cmd cron.Job) (cron.EntryID, error) {
	return crontab.Cron.AddJob(spec, cmd)
}

func (crontab *Crontab) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	return crontab.Cron.AddFunc(spec, func() {
		crontab.middlewareDispatcher(cmd)
	})
}

func (crontab *Crontab) Run() {
	crontab.Cron.Start()
}

func (crontab *Crontab) AddMiddleware(task Cron) {
	ml := len(crontab.middlewares)
	if ml == 0 {
		crontab.middlewares = []Cron{task}
	} else {
		nextTask := crontab.middlewares[ml-1]
		crontab.middlewares = append(crontab.middlewares, func(next func()) {
			task(func() {
				nextTask(next)
			})
		})
	}
}

func (crontab *Crontab) middlewareDispatcher(next func()) {
	ml := len(crontab.middlewares)
	if ml == 0 {
		next()
	} else {
		crontab.middlewares[ml-1](next)
	}
}

func (crontab *Crontab) Exit() {
	crontab.exit.Do(func() {
		crontab.Cron.Stop()
	})
}
