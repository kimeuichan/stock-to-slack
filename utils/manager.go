package utils

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/domain"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"github.com/spf13/viper"
	"time"
)

type StockManager struct {
	Scheduler   *gocron.Scheduler
	StockClient client.StockAsyncClient
	StockSender sender.SendClient
	Interval    uint64
	stocks      map[string]bool
}

var defaultTag = []string{"STOCK"}

func NewStockManager() *StockManager {
	scheduler := gocron.NewScheduler(time.Local)
	return &StockManager{Scheduler: scheduler, stocks: make(map[string]bool)}
}

func (sm *StockManager) SubscribeStocks(stockNumbers []string){
	for _, stockNumber := range stockNumbers {
		sm.SubscribeStock(stockNumber)
	}
}

func (sm *StockManager) SubscribeStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; exists {
		return
	}

	tempTag := append(defaultTag, stockNumber)
	sm.Scheduler.SetTag(tempTag).Every(sm.Interval).Second().Do(func() {
		stockSummary := make(chan domain.StockSummary)
		errChannel := make(chan error)
		go sm.StockClient.GetStockSummaryByGoRoutine(viper.GetString("STOCK_NUMBER"), stockSummary, errChannel)

		select {
		case stock := <-stockSummary:
			sm.stocks[stockNumber] = true

			if err := sm.StockSender.SendStock(&stock); err != nil {
				panic(err)
			}
		case err := <-errChannel:
			panic(err)
		}
	})
}

func (sm *StockManager) DetachStock(stockNumber string) {
	if _, exists := sm.stocks[stockNumber]; !exists {
		return
	}

	sm.Scheduler.RemoveJobByTag(stockNumber)
	delete(sm.stocks, stockNumber)
}

func (sm *StockManager) Clear() {
	for k := range sm.stocks {
		sm.Scheduler.RemoveJobByTag(k)
	}
}
