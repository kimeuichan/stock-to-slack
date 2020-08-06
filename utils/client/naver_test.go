package client

import (
	"github.com/kimeuichan/stock-to-slack/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/korean"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestNaverClient_GetStockSummary(t *testing.T) {
	jsonString := "{\"item_list\":[{\"itemcode\":\"019170\",\"itemname\":\"신풍제약\",\"change_val\":\"1,900\",\"change_rate\":\"-2.59\",\"now_val\":\"71,600\",\"risefall\":\"5\"}],\"prev_page\":0,\"next_page\":0,\"itemTotalCnt\":1,\"login\":\"false\",\"type\":\"recent\",\"page\":1,\"code\":\"019170\",\"sel_cid\":null}"
	encodeString, _ := korean.EUCKR.NewEncoder().String(jsonString)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(encodeString))
	}))

	defer server.Close()

	naverClient := NaverClient{Host: server.URL, client: http.DefaultClient}

	stockSummary, _ := naverClient.GetStockSummary("019170")
	expectedStockSummary := &utils.StockSummary{
		ChangeVal:  "1,900",
		ChangeRate: "-2.59",
		StockName:  "신풍제약",
		NowVal:     "71,600",
	}

	assert.Equal(t, expectedStockSummary, stockSummary)
}

func TestNaverClient_GetStockSummaryByGoRoutine(t *testing.T) {
	jsonString := "{\"item_list\":[{\"itemcode\":\"019170\",\"itemname\":\"신풍제약\",\"change_val\":\"1,900\",\"change_rate\":\"-2.59\",\"now_val\":\"71,600\",\"risefall\":\"5\"}],\"prev_page\":0,\"next_page\":0,\"itemTotalCnt\":1,\"login\":\"false\",\"type\":\"recent\",\"page\":1,\"code\":\"019170\",\"sel_cid\":null}"
	encodeString, _ := korean.EUCKR.NewEncoder().String(jsonString)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(encodeString))
	}))

	defer server.Close()

	naverClient := NaverClient{Host: server.URL, client: http.DefaultClient}

	stockSummary := make(chan utils.StockSummary)
	err := make(chan error)

	go naverClient.GetStockSummaryByGoRoutine("019170", stockSummary, err)
	expectedStockSummary := utils.StockSummary{
		ChangeVal:  "1,900",
		ChangeRate: "-2.59",
		StockName:  "신풍제약",
		NowVal:     "71,600",
	}

	assert.Equal(t, expectedStockSummary, <-stockSummary)
}
