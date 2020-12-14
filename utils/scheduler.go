package utils

import (
	"github.com/go-co-op/gocron"
	"time"
)

type StockScheduler struct {
	Scheduler   *gocron.Scheduler
}

func NewStockScheduler() *StockScheduler {
	return &StockScheduler{Scheduler: gocron.NewScheduler(time.Local)}
}

