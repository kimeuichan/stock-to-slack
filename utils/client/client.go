package client

import (
	"github.com/kimeuichan/stock-to-slack/utils"
	"net/http"
)

type StockClient interface {
	GetStockSummary(stockNumber string) (*utils.StockSummary, error)
}

func GetClient(clientType string) StockClient {
	var client StockClient = nil

	if clientType == "naver" {
		tempClient := new(NaverClient)
		tempClient.host = NaverStockURI
		tempClient.client = new(http.Client)
		tempClient.client.Transport = NaverHeader{r: http.DefaultTransport}
		client = tempClient
	}

	return client
}
