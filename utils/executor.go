package utils

import (
	"github.com/kimeuichan/stock-to-slack/utils/client"
	"github.com/kimeuichan/stock-to-slack/utils/sender"
)

type StockExecutor struct {
	stockClient client.StockAsyncClient
	senderClient sender.SendClient
}

func NewStockExecutor(stockClient client.StockAsyncClient, senderClient sender.SendClient) *StockExecutor {
	return &StockExecutor{stockClient: stockClient, senderClient: senderClient}
}

func (se *StockExecutor) Execute(stockNumbers []string)  {
	stocksChannel := se.stockClient.GetStockSummaryByGoRoutine(stockNumbers)
	se.senderClient.SendStocks(stocksChannel)
}





