package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const NaverStockURI = "https://finance.naver.com/"

type StockSummary struct {
	ChangeVal  string
	ChangeRate string
	StockName  string
	NowVal     string
}

type StockClient interface {
	GetStockSummary(stockNumber string) (*StockSummary, error)
}

type NaverClient struct {
	host   string
	client *http.Client
}

type NaverHeader struct {
	r http.RoundTripper
}

type naverItem struct {
	ItemCode   string `json:"itemcode"`
	ItemName   string `json:"itemname"`
	ChangeVal  string `json:"change_val"`
	ChangeRate string `json:"change_rate"`
	NowVal     string `json:"now_val"`
	RiseFall   string `json:"risefall"`
}

type NaverStockResponse struct {
	ItemList     []naverItem `json:"item_list"`
	PrevPage     int         `json:"prev_page"`
	NextPage     int         `json:"prev_page"`
	ItemToTalCnt int         `json:"itemTotalCnt"`
	Login        bool        `json:"login"`
	ReqType      string      `json:"type"`
	Page         int         `json:"page"`
	Code         string      `json:"string"`
}

func (nh NaverHeader) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("authority", "finance.naver.com")
	r.Header.Add("content-length", "0")
	r.Header.Add("pragma", "no-cache")
	r.Header.Add("cache-control", "no-cache")
	r.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	r.Header.Add("accept", "*/*")
	r.Header.Add("origin", "https://finance.naver.com")
	r.Header.Add("sec-fetch-site", "same-origin")
	r.Header.Add("sec-fetch-mode", "cors")
	r.Header.Add("sec-fetch-dest", "empty")
	r.Header.Add("accept-language", "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6")
	r.Header.Add("cookie", "naver_stock_codeList=019170%7C;")
	return nh.r.RoundTrip(r)
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

func (nc *NaverClient) Get(url string) (resp *http.Response, err error) {
	return nc.client.Get(nc.host + url)
}

func (nc *NaverClient) GetStockSummary(stockNumber string) (*StockSummary, error) {
	naverStockFullUrl := "/item/item_right_ajax.nhn?type=recent&code=" + stockNumber + "&page=1"

	resp, err := nc.Get(naverStockFullUrl)

	if err != nil {
		return nil, err
	}

	var naverStock NaverStockResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&naverStock)

	if err != nil {
		return nil, err
	}

	stockInfo := naverStock.ItemList[0]
	fmt.Println(stockInfo)

	stockSummary := &StockSummary{
		ChangeVal:  stockInfo.ChangeVal,
		ChangeRate: stockInfo.ChangeRate,
		StockName:  stockInfo.ItemName,
		NowVal:     stockInfo.NowVal,
	}

	return stockSummary, err
}
