package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kimeuichan/stock-to-slack/domain"
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

var interval uint64
const BaseStockTask = "BASE_STOCK_TASK"

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

	baseStockTask := func() {
		stockExecutor.Execute(stockManager.GetStocks())
	}

	stockScheduler.Scheduler.Every(interval).Seconds().SetTag([]string{BaseStockTask}).Do(baseStockTask)

	stockScheduler.Scheduler.StartAsync()

	r := gin.Default()

	r.POST("/stocks", func(c *gin.Context) {
		var stockRequest domain.StockRequest
		if err := c.BindJSON(&stockRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		stockManager.AttachStock(stockRequest.StockNumber)

		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.DELETE("/stocks", func(c *gin.Context) {
		var stockRequest domain.StockRequest
		if err := c.BindJSON(&stockRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		stockManager.DetachStock(stockRequest.StockNumber)

		c.JSON(200, gin.H{
			"message": "ok",
		})

	})

	r.POST("/interval", func(c *gin.Context) {
		var intervalRequest domain.IntervalRequest
		if err := c.BindJSON(&intervalRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		interval = intervalRequest.Interval
		stockScheduler.Scheduler.RemoveJobByTag(BaseStockTask)
		stockScheduler.Scheduler.Every(interval).Seconds().SetTag([]string{BaseStockTask}).Do(baseStockTask)

		c.JSON(200, gin.H{
			"message": "ok",
		})

	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
