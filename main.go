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

	stockManager := utils.NewStockManager()

	stockWorker := utils.NewStockWorker(nc, sc, viper.GetUint64("INTERVAL"))

	stockManager.Subscribe(stockWorker)

	stocks := strings.Split(viper.GetString("STOCK_NUMBERS"), ",")

	for _, stock := range stocks {
		stockManager.AttachStock(stock)
	}

	stockWorker.Scheduler.StartBlocking()
}
