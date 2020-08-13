package utils

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"time"
)

type StockWorker struct {
	Scheduler   *gocron.Scheduler
	StockClient client.StockAsyncClient
	StockSender sender.SendClient
	Interval    uint64
	stocks      map[string]bool
}

var defaultTag = []string{"STOCK"}

func NewStockWorker() *StockWorker {
	scheduler := gocron.NewScheduler(time.Local)
	return &StockWorker{Scheduler: scheduler, stocks: make(map[string]bool)}
}

func (sm *StockWorker) AttachStocks(stockNumbers []string){
	for _, stockNumber := range stockNumbers {
		sm.AttachStock(stockNumber)
	}
}

func (sm *StockWorker) AttachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; exists {
		return
	}

	tempTag := append(defaultTag, stockNumber)
	sm.Scheduler.Every(sm.Interval).Seconds().SetTag(tempTag).Do(func() {
		if stockSummary, err := sm.StockClient.GetStockSummary(stockNumber); err == nil {
			sm.StockSender.SendStock(stockSummary)
		} else {
			panic(err)
		}
	})
}

func (sm *StockWorker) DetachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; !exists {
		return
	}

	sm.Scheduler.RemoveJobByTag(stockNumber)
	delete(sm.stocks, stockNumber)
}

func (sm *StockWorker) ExpiredAllStocks() {
	sm.Scheduler.RemoveJobByTag(defaultTag[0])
}

func (sm *StockWorker) RecoverStocks() {
	for k := range sm.stocks {
		sm.AttachStock(k)
	}
}
