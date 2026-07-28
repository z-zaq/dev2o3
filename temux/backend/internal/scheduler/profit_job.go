package scheduler

import (
	"temux/internal/repository"
	"temux/internal/services"
	"time"
)

type ProfitJob struct {
	InvestmentRepo *repository.InvestmentRepository
}

func (j *ProfitJob) Name() string {
	return "profit-processor"
}

func (j *ProfitJob) Interval() time.Duration {
	return time.Minute
}

func (j *ProfitJob) Run() error {
	services.ProcessProfits(
		j.InvestmentRepo,
	)
	return nil
}
