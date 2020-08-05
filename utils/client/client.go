package client

import (
	"crypto/tls"
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
		tempClient.Host = NaverStockURI
		tempClient.client = new(http.Client)
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tempClient.client.Transport = NaverHeader{r: http.DefaultTransport}
		client = tempClient
	}

	return client
}
