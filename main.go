package main

import (
	"fmt"
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()

	nc := utils.GetClient("naver")
	sc := utils.NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))

	stockSummary, err := nc.GetStockSummary("019170")

	if err != nil {
		fmt.Println(stockSummary)
	}

	if err := sc.SendStock(stockSummary); err != nil {
		fmt.Println(err)
	}
}
