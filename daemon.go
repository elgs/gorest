// daemon
package gorest

import (
	"fmt"
	"github.com/elgs/cron"
)

type Job struct {
	MakeAction func(dbo DataOperator, r *Gorest) func()
	Cron       string
	Handler    int
}

var Sched *cron.Cron
var JobRegistry = make(map[string]*Job)

func RegisterJob(id string, job *Job) {
	JobRegistry[id] = job
}

func GetJob(id string) *Job {
	return JobRegistry[id]
}

func StartDaemons(dbo DataOperator, r *Gorest) {
	Sched = cron.New()
	for _, job := range JobRegistry {
		h, err := Sched.AddFunc(job.Cron, job.MakeAction(dbo, r))
		if err != nil {
			fmt.Println(err)
			continue
		}
		job.Handler = h
	}
	Sched.Start()
}
