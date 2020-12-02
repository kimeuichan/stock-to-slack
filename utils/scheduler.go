package utils

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
)

type StockScheduler struct {
	Scheduler   *gocron.Scheduler
	StockClient client.StockAsyncClient
	StockSender sender.SendClient
	StockManager *StockManager
	Interval    uint64
}

func NewStockScheduler(stockClient client.StockAsyncClient, stockSender sender.SendClient, stockManager *StockManager, interval uint64) *StockScheduler {
	return &StockScheduler{StockClient: stockClient, StockSender: stockSender, StockManager: stockManager, Interval: interval}
}

var defaultTag = []string{"STOCK"}


func (sw *StockScheduler) Execute(){
	stocksChannel := sw.StockClient.GetStockSummaryByGoRoutine(sw.StockManager.GetStocks())
	sw.StockSender.SendStocks(stocksChannel)
}

func (sw *StockScheduler) AttachStock(stockNumber string) {
	tempTag := append(defaultTag, stockNumber)
	sw.Scheduler.Every(sw.Interval).Seconds().SetTag(tempTag).Do(func() {
		sw.Execute()
	})
}

func (sw *StockScheduler) DetachStock(stockNumber string) {
	sw.Scheduler.RemoveJobByTag(stockNumber)
}

func (sw *StockScheduler) ExpiredAllStocks() {
	sw.Scheduler.RemoveJobByTag(defaultTag[0])
}
