package main

import (
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

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
