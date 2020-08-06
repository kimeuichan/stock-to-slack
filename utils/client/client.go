package client

import (
	"crypto/tls"
	"github.com/kimeuichan/stock-to-slack/utils"
	"net/http"
)

type StockClient interface {
	GetStockSummary(stockNumber string) (*utils.StockSummary, error)
}

type StockAsyncClient interface {
	StockClient
	GetStockSummaryByGoRoutine(stockNumber string, c chan utils.StockSummary, err chan error)
}

func GetClient(clientType string) StockAsyncClient {
	var client StockAsyncClient = nil

	if clientType == "naver" {
		tempClient := new(NaverClient)
		tempClient.Host = NaverStockURI
		tempClient.client = new(http.Client)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tempClient.client.Transport = NaverHeader{r: http.DefaultTransport}
		client = tempClient
	}

	return client
}
