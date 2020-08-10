package main

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/domain"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"github.com/spf13/viper"
	"time"
)

func main() {
	viper.AutomaticEnv()
	s := gocron.NewScheduler(time.Local)

	nc := client.GetClient("naver")
	sc := sender.NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))

	getStock := func() {
		stockSummary := make(chan domain.StockSummary)
		errChannel := make(chan error)
		go nc.GetStockSummaryByGoRoutine(viper.GetString("STOCK_NUMBER"), stockSummary, errChannel)

		select {
		case stock := <-stockSummary:
			if err := sc.SendStock(&stock); err != nil {
				panic(err)
			}
		case err := <-errChannel:
			panic(err)
		}
	}

	s.Every(1).Minutes().StartImmediately().Do(getStock)
	s.StartBlocking()
}
