package utils

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"time"
)

type StockManager struct {
	Scheduler *gocron.Scheduler
	StockClient *client.StockClient
}

var defaultTag = []string{"STOCK"}

func NewStockManager() *StockManager {
	scheduler := gocron.NewScheduler(time.Local)
	return &StockManager{Scheduler: scheduler}
}

func (sm *StockManager) AddStock(stockNumber string) {
	tempTag := append(defaultTag, stockNumber)
	sm.Scheduler.SetTag(tempTag).Every(1).Minutes().Do(func() {

	})
}



