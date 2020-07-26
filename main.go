package main

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/spf13/viper"
	"strings"
	"time"
)

func main() {
	viper.AutomaticEnv()
	s := gocron.NewScheduler(time.Local)

	nc := client.GetClient("naver")
	sc := utils.NewSlack(viper.GetString("SLACK_WEBHOOK_URL"))

	getStock := func(stockNumber string) {
		fmt.Println(stockNumber)
		stockSummary, err := nc.GetStockSummary(stockNumber)

		if err != nil {
			panic(err)
		}

		if err := sc.SendStock(stockSummary); err != nil {
			panic(err)
		}
	}

	stocks := viper.GetString("STOCK_NUMBER")

	for _, v := range strings.Split(stocks, ","){
		s.Every(1).Minutes().SetTag("STOCK").Do(getStock, v)
	}

	exfireStock := func(){
		s.RemoveJobByTag("STOCK")
	}

	s.StartAt().Do(exfireStock)

	s.StartBlocking()
}
