package scheduler

import (
	"log"
	"time"
)

type Job interface {
	Name() string
	Run() error
	Interval() time.Duration
}

type Scheduler struct {
	jobs []Job
}

func New() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) Register(job Job) {
	s.jobs = append(s.jobs, job)
}

func (s *Scheduler) Start() {
	for _, job := range s.jobs {
		go s.run(job)
	}
}

func (s *Scheduler) run(job Job) {
	ticker := time.NewTicker(job.Interval())
	defer ticker.Stop()

	log.Printf("[scheduler] started job: %s", job.Name())

	for {
		if err := job.Run(); err != nil {
			log.Printf(
				"[scheduler] %s failed: %v",
				job.Name(),
				err,
			)
		}

		<-ticker.C
	}
}
