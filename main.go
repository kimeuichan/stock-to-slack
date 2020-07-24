package main

import (
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/spf13/viper"
	"time"
)

func main() {
	viper.AutomaticEnv()
	s := gocron.NewScheduler(time.Local)

	getStock := func() {
		nc := client.GetClient("naver")
		sc := utils.NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))

		stockSummary, err := nc.GetStockSummary(viper.GetString("STOCK_NUMBER"))

		if err != nil {
			panic(err)
		}

		if err := sc.SendStock(stockSummary); err != nil {
			panic(err)
		}
	}

	s.Every(1).Minutes().Do(getStock)
	s.StartBlocking()
}
