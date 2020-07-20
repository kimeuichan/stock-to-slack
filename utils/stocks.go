package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type StockSummary struct {
	changeVal string
	changeRate string
	stockName string
	nowVal string
}


type StockClient interface {
	GetStockSummary(int) (*StockSummary,  error)
}

type NaverClient struct {
	client *http.Client
	header *http.Header
}


type NaverHeader struct {
	r http.RoundTripper
}

func (nh NaverHeader) RoundTrip(r *http.Request) (*http.Response, error){
	r.Header.Add("authority", "finance.naver.com")
	r.Header.Add("content-length", "0")
	r.Header.Add("pragma", "no-cache")
	r.Header.Add("cache-control", "no-cache")
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	r.Header.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	r.Header.Add("accept", "*/*")
	r.Header.Add("origin", "https://finance.naver.com")
	r.Header.Add("sec-fetch-site", "same-origin")
	r.Header.Add("sec-fetch-mode", "cors")
	r.Header.Add("sec-fetch-dest", "empty")
	r.Header.Add("accept-language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")
	r.Header.Add("cookie", "naver_stock_codeList=019170%7C;")
	return nh.r.RoundTrip(r)
}

func GetClient(clientType string) *StockClient {
	var client *StockClient

	if clientType == "naver" {
		client = new(NaverClient)
		client.client = new (http.Client)
		client.client.Transport = NaverHeader{r: http.DefaultTransport}
	}

	return nil
}

func (nc *NaverClient) Get (url string) (resp *http.Response, err error){
	return nc.client.Get(url)
}


func (nc *NaverClient) GetStockSummary(stockNumber int) (*StockSummary,  error){
	naverStockFullUrl := "?type=recent&code=" + string(stockNumber) + "&page=1"

	resp, err := nc.Get(naverStockFullUrl)

	if err != nil{
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil, err
}