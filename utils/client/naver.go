package client

import (
	"encoding/json"
	"fmt"
	"github.com/kimeuichan/stock-to-slack/domain"
	"golang.org/x/text/encoding/korean"
	"io"
	"net/http"
	"sync"
)

const NaverStockURI = "https://finance.naver.com"

type NaverClient struct {
	Host   string
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
	Login        string      `json:"login"`
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
	return nh.r.RoundTrip(r)
}

func (nc *NaverClient) GetStockSummary(stockNumber string) (*domain.StockSummary, error) {
	naverStockFullUrl := "/item/item_right_ajax.nhn?type=recent&code=" + stockNumber + "&page=1"

	request, err := http.NewRequest("GET", nc.Host+naverStockFullUrl, nil)
	request.Header.Add("cookie", fmt.Sprintf("naver_stock_codeList=%s;", stockNumber))

	resp, err := nc.client.Do(request)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var naverStock NaverStockResponse

	tempByte := make([]byte, 1000)

	n, err := korean.EUCKR.NewDecoder().Reader(resp.Body).Read(tempByte)

	if err != nil {
		if err != io.EOF {
			return nil, err
		}
	}

	if err = json.Unmarshal(tempByte[:n], &naverStock); err != nil {
		return nil, err
	}

	stockInfo := naverStock.ItemList[0]

	stockSummary := &domain.StockSummary{
		ChangeVal:  stockInfo.ChangeVal,
		ChangeRate: stockInfo.ChangeRate,
		StockName:  stockInfo.ItemName,
		NowVal:     stockInfo.NowVal,
	}

	return stockSummary, err
}

func (nc *NaverClient) GetStockSummaryByGoRoutine(stockNumbers []string) chan *domain.StockSummary {
	out := make(chan *domain.StockSummary, len(stockNumbers))
	var wg sync.WaitGroup

	wg.Add(len(stockNumbers))

	for _, stock := range stockNumbers {
		go func() {
			// TODO: error handling
			if stockSummary, tempErr := nc.GetStockSummary(stock); tempErr == nil {
				out <- stockSummary
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
