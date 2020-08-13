package main

import (
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	viper.AutomaticEnv()

	nc := client.GetClient("naver")
	sc := sender.NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))

	manager := utils.NewStockWorker()

	manager.StockClient = nc
	manager.StockSender = sc
	manager.Interval = viper.GetUint64("INTERVAL")

	stocks := strings.Split(viper.GetString("STOCK_NUMBERS"), ",")

	manager.AttachStocks(stocks)

	manager.Scheduler.StartBlocking()

}
