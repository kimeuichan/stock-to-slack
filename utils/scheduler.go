package utils

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"time"
)

type StockScheduler struct {
	Scheduler   *gocron.Scheduler
	StockClient client.StockClient
	StockSender sender.SendClient
	Interval    uint64
}

var defaultTag = []string{"STOCK"}

func NewStockWorker(stockClient client.StockClient, sendClient sender.SendClient, interval uint64) *StockScheduler {
	scheduler := gocron.NewScheduler(time.Local)
	return &StockScheduler{Scheduler: scheduler, StockClient: stockClient, StockSender: sendClient, Interval: interval}
}

func (sw *StockScheduler) AttachStock(stockNumber string) {
	tempTag := append(defaultTag, stockNumber)
	sw.Scheduler.Every(sw.Interval).Seconds().SetTag(tempTag).Do(func() {
		if stockSummary, err := sw.StockClient.GetStockSummary(stockNumber); err == nil {
			sw.StockSender.SendStock(stockSummary)
		} else {
			panic(err)
		}
	})
}

func (sw *StockScheduler) DetachStock(stockNumber string) {
	sw.Scheduler.RemoveJobByTag(stockNumber)
}

func (sw *StockScheduler) ExpiredAllStocks() {
	sw.Scheduler.RemoveJobByTag(defaultTag[0])
}
