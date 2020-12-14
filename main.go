package main

import (
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"github.com/spf13/viper"
	"strings"
)

var interval uint64

func init() {
	viper.AutomaticEnv()
	interval = viper.GetUint64("INTERVAL")
}

func main() {
	nc := client.GetClient("naver")
	sc := sender.GetSender(viper.GetString("SENDER"))
	stockExecutor := utils.NewStockExecutor(nc, sc)

	stockManager := utils.NewStockManager()
	stocks := strings.Split(viper.GetString("STOCK_NUMBERS"), ",")
	for _, stock := range stocks {
		stockManager.AttachStock(stock)
	}

	stockScheduler := utils.NewStockScheduler()

	stockScheduler.Scheduler.Every(interval).Seconds().SetTag([]string{"test"}).Do(func() {
		stockExecutor.Execute(stockManager.GetStocks())
	})

	stockScheduler.Scheduler.StartBlocking()
}
