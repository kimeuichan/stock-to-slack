package client

import (
	"crypto/tls"
	"github.com/kimeuichan/stock-to-slack/domain"
	"net/http"
)

type StockClient interface {
	GetStockSummary(stockNumber string) (*domain.StockSummary, error)
}

type StockAsyncClient interface {
	StockClient
	GetStockSummaryByGoRoutine(stockNumber string, c chan domain.StockSummary, err chan error)
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
